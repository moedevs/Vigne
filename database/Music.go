package database

import (
	"github.com/bela333/Vigne/errors"
)

type MusicDatabase struct {
	d *Database
}

func (d *Database) Music() (*MusicDatabase, error) {
	if !d.Redis.HExists(d.Decorate("config"), "musicChannel").Val() {
		return nil, errors.NoMusic
	}
	if !d.Redis.HExists(d.Decorate("config"), "musicVoiceChannel").Val() {
		return nil, errors.NoMusic
	}
	return &MusicDatabase{d}, nil
}

func (d MusicDatabase) GetChannel() string {
	return d.d.Redis.HGet(d.d.Decorate("config"), "musicChannel").Val()
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
	return d.d.Redis.HGet(d.d.Decorate("config"), "musicVoiceChannel").Val()
}