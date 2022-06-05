package service

import (
	"errors"
	"strconv"
	"strings"
)

//正常的body长度在16万左右,这里小于1000认为访问失败
const MinBodyLength = 1000

var CmsMap CompactMsMap

type CompactMsMap map[int64]*CompactMs

type LanguageMsg struct {
	// Subtitles languages of subtitles, 字幕语言
	Subtitles []string
	// Audios languages of audios, 音频语言
	Audios []string
}

type CompactMs struct {
	// Id primary key, 主键      即url中路径 /title/ 后面的数字
	Id int64
	// Title title in English, 英文标题
	Title string
	// Regions available regions, 可观看地区
	Regions []string
	//地区对语言是一对多的关系
	Region2Language map[string]LanguageMsg
}

func InitCmsMap() {
	CmsMap = make(CompactMsMap)
}

// ParseDetail parses one movie subject metadata
func ParseDetail(CmsMap CompactMsMap, body []byte) (rs *CompactMs, err error) {
	//<p class="footer-country">Netflix Singapore</p>    					   定位country
	//<h1 class="title-title" data-uia="title-info-title">Strong Island</h1>   定位英文标题
	//<meta property="og:url" content="https://www.netflix.com/sg/title/80168230" id="meta-url" />   定位Id
	//data-uia="more-details-item-audio">English - Audio Description			定位Audio
	//<!-- -->,
	//</span><span class="more-details-item item-audio"
	//data-uia="more-details-item-audio">English [Original]
	//<!-- -->,
	//</span><span class="more-details-item item-audio"
	//data-uia="more-details-item-audio">Japanese</span></div>

	//data-uia="more-details-item-subtitle">English								定位language
	//<!-- -->,
	//</span><span class="more-details-item item-subtitle"
	//data-uia="more-details-item-subtitle">Simplified Chinese
	//<!-- -->,
	//</span><span class="more-details-item item-subtitle"
	//data-uia="more-details-item-subtitle">Traditional Chinese</span></div>

	mes := string(body)
	if len(body) < MinBodyLength {
		return nil, errors.New("Invalid Body ")
	}

	//定位url
	i := strings.Index(mes, "<meta property=\"og:url\" content=\"")
	j := strings.Index(mes[i:], "\" id=\"meta-url\"")
	url := mes[i+len("<meta property=\"og:url\" content=\"") : i+j]

	//定位id
	i = strings.Index(url, "/title/")
	s := url[i+len("/title/"):]
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, errors.New("Wrong Url Type ")
	}

	cms, exist := CmsMap[id]
	if !exist {
		cms = &CompactMs{
			Id:              id,
			Regions:         []string{},
			Region2Language: make(map[string]LanguageMsg),
		}
		CmsMap[id] = cms

		//Cms 标题
		i = strings.Index(mes, "<h1 class=\"title-title\" data-uia=\"title-info-title\">")
		j = strings.Index(mes[i:], "</h1>")
		cms.Title = mes[i+len("<h1 class=\"title-title\" data-uia=\"title-info-title\">") : i+j]
	}

	//cms 增加地区
	i = strings.Index(mes, "<p class=\"footer-country\">Netflix")
	j = strings.Index(mes[i:], "</p>")
	region := mes[i+len("<p class=\"footer-country\">Netflix ") : i+j]
	cms.Regions = append(cms.Regions, region)

	//获取Cms 的language信息
	var langMsg = LanguageMsg{
		[]string{},
		[]string{},
	}

	findAllAudio := false
	m1 := mes
	for !findAllAudio { //循环查找,直到找到所有的Audio
		i = strings.Index(m1, "data-uia=\"more-details-item-audio\">")
		j = strings.Index(m1[i:], "<")
		audio := m1[i+len("data-uia=\"more-details-item-audio\">") : i+j]
		langMsg.Audios = append(langMsg.Audios, audio)

		m1 = m1[i+len("data-uia=\"more-details-item-audio\">"):]
		if strings.Index(m1, "data-uia=\"more-details-item-audio\">") == -1 {
			findAllAudio = true
		}
	}

	findAllSubtitles := false
	m2 := mes
	for !findAllSubtitles {
		i = strings.Index(m2, "data-uia=\"more-details-item-subtitle\">")
		j = strings.Index(m2[i:], "<")
		subtitle := m2[i+len("data-uia=\"more-details-item-subtitle\">") : i+j]
		langMsg.Subtitles = append(langMsg.Subtitles, subtitle)

		m2 = m2[i+len("data-uia=\"more-details-item-subtitle\">"):]
		if strings.Index(m2, "data-uia=\"more-details-item-subtitle\">") == -1 {
			findAllSubtitles = true
		}
	}
	//匹配地区和语言
	cms.Region2Language[region] = langMsg

	return cms, nil
}
