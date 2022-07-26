package solaredge

import (
	"fmt"
	"strings"
)

type SitesEnergy struct {
	TimeUnit       string `json:"timeUnit"`
	Unit           string `json:"unit"`
	Count          int64  `json:"count"`
	SiteEnergyList []struct {
		SiteId       int64 `json:"siteId"`
		EnergyValues struct {
			MeasuredBy string            `json:"measuredBy"`
			Values     []SiteEnergyValue `json:"values"`
		} `json:"energyValues"`
	} `json:"siteEnergyList"`
}

type SitesEnergyResponse struct {
	Energy SitesEnergy `json:"sitesEnergy"`
}

func (s *SitesService) Energy(siteIDS []int64, energyOptions *SiteEnergyRequest) (SitesEnergy, error) {
	IDS := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(siteIDS)), ","), "[]")
	u, err := addOptions(fmt.Sprintf("/sites/%s/energy", IDS), energyOptions)
	if err != nil {
		return SitesEnergy{}, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return SitesEnergy{}, err
	}
	var siteEnergyResponse SitesEnergyResponse
	_, err = s.client.do(req, &siteEnergyResponse)
	return siteEnergyResponse.Energy, err
}

// This is broken
func (s *SitesService) TimeFrameEnergy(siteIDS []int64, energyOptions *SiteEnergyRequest) (SitesEnergy, error) {
	IDS := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(siteIDS)), ","), "[]")
	u, err := addOptions(fmt.Sprintf("/sites/%s/timeFrameEnergy/", IDS), energyOptions)
	if err != nil {
		return SitesEnergy{}, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return SitesEnergy{}, err
	}
	var siteEnergyResponse SitesEnergyResponse
	_, err = s.client.do(req, &siteEnergyResponse)
	return siteEnergyResponse.Energy, err
}
