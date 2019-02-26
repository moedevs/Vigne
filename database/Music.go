package database

import (
	"github.com/moedevs/Vigne/errors"
	"strconv"
	"time"
)

type MusicDatabase struct {
	//TODO: Remove. Legacy, for most uses except for config
	d *Database
	MusicChannel StringValue
	MusicVoiceChannel StringValue
	PlayLive StringValue
	MaxDuration StringValue
}

func (d *Database) Music() (*MusicDatabase, error) {
	cfgHandler := d.NewConfigHandler()
	m := MusicDatabase{}
	m.d = d
	var exists bool
	m.MusicChannel, exists = cfgHandler.OptionalValue("musicChannel")
	if !exists {
		return nil, errors.NoMusic
	}
	m.MusicVoiceChannel, exists = cfgHandler.OptionalValue("musicVoiceChannel")
	if !exists {
		return nil, errors.NoMusic
	}
	m.PlayLive, _ = cfgHandler.OptionalValue("canPlayLive")
	m.MaxDuration, _ = cfgHandler.OptionalValue("maxMusicDuration")
	return &m, nil
}

func (d MusicDatabase) GetChannel() string {
	return d.MusicChannel.Get()
}

func (d MusicDatabase) PopNext() string {
	return d.d.Redis.LPop(d.d.Decorate("musicQueue")).Val()
}

func (d MusicDatabase) AddSong(data []byte) error {
	return d.d.Redis.RPush(d.d.Decorate("musicQueue"), string(data)).Err()
}

func (d MusicDatabase) IsValidExtractor(extractor string) bool {
	if d.d.Redis.Exists(d.d.Decorate("validExtractors")).Val() == 0 {
		//If validExtractors doesn't exist, we accept the extractor anyway
		return true
	}
	return d.d.Redis.SIsMember(d.d.Decorate("validExtractors"), extractor).Val()
}

func (d MusicDatabase) GetVoiceChannel() string {
	return d.MusicVoiceChannel.Get()
}

func (d MusicDatabase) CanPlay(duration time.Duration) bool {
	if d.MaxDuration.Get() == "" {
		//No maxMusicDuration is set
		return true
	}
	max, err := strconv.Atoi(d.MaxDuration.Get())
	if err != nil {
		//Couldn't get maxMusicDuration
		return true
	}
	if time.Duration(max)*time.Second < duration {
		//music duration is larger than maxMusicDuration. Don't play it.
		return false
	}
	return true

}

func (d MusicDatabase) CanPlayLive() bool {
	if d.PlayLive.Get() == ""{
		return true
	}
	val, err := strconv.Atoi(d.PlayLive.Get())
	if err != nil {
		return  true
	}
	if val != 0 {
		return true
	}

	return false
}