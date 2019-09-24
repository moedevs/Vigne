package levels

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/database/interfaces"
	"github.com/moedevs/Vigne/server"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type LevelsModule struct {
	LevelContainer interfaces.IntegerMapValue
	ThresholdContainer interfaces.MapValue
	mutex sync.Mutex
	userSet []string
	threadRunning bool
	s *server.Server
}

const(
	levelBoundMin = 25
	levelBoundMax = 50
	levelUp = 300	//On average about 1 level up for every 8 minutes
)

func (m LevelsModule) GetName() string {
	return "levels"
}

func (m *LevelsModule) Init(s *server.Server) error {
	m.s = s
	s.Session.AddHandler(m.onMessage)
	m.LevelContainer = s.Database.IntegerMap("levels")
	m.ThresholdContainer = s.Database.Map("level_threshold")
	m.userSet = []string{}

	//Get command handler module
	cmdInterface, err := s.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := (cmdInterface).(*commands.CommandsModule)
	err = cmd.RegisterCommand(&LevelCommand{m})
	if err != nil {
		return err
	}

	return nil
}

func (m *LevelsModule) GetUserCounter(userid string) interfaces.IntegerValue {
	return m.LevelContainer.Get(userid)
}

type sortableValue struct {
	value int
	payload string
}

type sortable []sortableValue

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	return s[i].value < s[j].value
}

func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func mapToSortable(m map[string]interfaces.StringValue) (sortable, error) {
	sorted := make(sortable, len(m))
	i := 0
	for key, val := range m {
		keyInt, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		sorted[i] = sortableValue{
			value:   keyInt,
			payload: val.Get(),
		}
		i++
	}
	sort.Sort(sorted)
	return sorted, nil
}

func (m *LevelsModule) onExpAdded(userid string, oldAmount int, guild string) error {
	current, err := m.GetExperience(userid)
	if err != nil {
		return err
	}
	thresholds, err:= m.ThresholdContainer.GetAll()
	if err == interfaces.ErrorNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	sorted, err := mapToSortable(thresholds)
	if err != nil {
		return err
	}

	for _, levelInfo := range sorted{
		if oldAmount < levelInfo.value*levelUp && current >= levelInfo.value*levelUp {
			//Found role that should be given out
			err = m.s.Session.GuildMemberRoleAdd(guild, userid, levelInfo.payload)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}

func (m *LevelsModule) GetLevels(userid string) (float64, error) {
	exp, err := m.GetExperience(userid)
	if err != nil {
		return 0, err
	}
	return float64(exp)/levelUp, nil
}

func (m *LevelsModule) GetExperience(userid string) (int, error) {
	return m.LevelContainer.Get(userid).Get()
}

func (m *LevelsModule) AddExperience(userid string, amount int, guild string) error {
	counter := m.GetUserCounter(userid)
	oldAmount, err := counter.Get()
	if err != nil {
		return err
	}
	err = counter.Add(amount)
	if err != nil {
		return err
	}
	err = m.onExpAdded(userid, oldAmount,guild)
	if err != nil {
		return err
	}
	return nil
}

func (m *LevelsModule) AddRandomExperience(userid string, guild string) error {
	levelAmount := rand.Intn(levelBoundMax-levelBoundMin+1)+levelBoundMin
	return m.AddExperience(userid, levelAmount, guild)
}

func (m *LevelsModule) counterThread(guild string)  {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		m.mutex.Lock()
		if len(m.userSet) <= 0 {
			ticker.Stop()
			m.threadRunning = false
			m.mutex.Unlock()
			return
		}
		var err error
		for _, user := range m.userSet {
			err = m.AddRandomExperience(user, guild)
			if err != nil {
				fmt.Printf("Couldn't give experience to %s: %s.\n", user, err)
				continue
			}
		}
		m.userSet = []string{}
		m.mutex.Unlock()
	}
}

func (m *LevelsModule) onMessage(s *discordgo.Session, e *discordgo.MessageCreate)  {
	if e.Author.Bot {
		//We don't serve droids here
		return
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	//Make sure that the user hasn't been added to the list already
	for _, user := range m.userSet {
		if user == e.Author.ID {
			return
		}
	}
	//Start thread is it isn't running already
	if !m.threadRunning{
		m.threadRunning = true
		go m.counterThread(e.GuildID)
	}
	m.userSet = append(m.userSet, e.Author.ID)
}

