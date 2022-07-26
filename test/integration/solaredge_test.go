package integration

import (
	"github.com/elliott-davis/solaredge-go/solaredge"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	client *solaredge.Client
	siteID int64
)

func init() {
	token := os.Getenv("SOLAREDGE_AUTH_TOKEN")
	id, err := strconv.ParseInt(os.Getenv("SOLAREDGE_SITE_ID"), 10, 64)
	if err != nil {
		panic("Can't get site ID")
	}
	siteID = id
	client = solaredge.NewClient(nil, token)
}

func TestSitesList(t *testing.T) {
	sites, err := client.Site.List(&solaredge.ListOptions{})
	if err != nil {
		t.Error(err)
	}
	if len(sites) == 0 {
		t.Error("no sites returned")
	}
	log.Printf("Found %d sites", len(sites))
}

func TestSiteDetails(t *testing.T) {
	site, err := client.Site.Details(siteID)
	if err != nil {
		t.Error(err)
	}
	if site.ID != siteID {
		t.Fatalf("site ID does not match")
	}
}

func TestSitesEnergies(t *testing.T) {
	ids := []int64{2249706, 1051426, 2935522}

	energies, err := client.Sites.Energy(ids, &solaredge.SiteEnergyRequest{
		TimePeriodRequest: solaredge.TimePeriodRequest{
			StartDate: solaredge.YMDTime(time.Date(2022, 07, 25, 0, 0, 0, 0, time.UTC)),
			EndDate:   solaredge.YMDTime(time.Date(2022, 07, 26, 0, 0, 0, 0, time.UTC)),
		},
		TimeUnit: "HOUR",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(energies.SiteEnergyList) == 0 {
		log.Fatalf("no sites returned")
	}

	for _, energy := range energies.SiteEnergyList {
		log.Println("Site ID:", energy.SiteId)
		log.Println("Measure By:", energy.EnergyValues.MeasuredBy)
		for _, value := range energy.EnergyValues.Values {
			if value.Value != nil {
				log.Printf("%s: %f", value.Date, *value.Value)
			} else {
				log.Printf("%s: ---", value.Date)
			}
		}
	}
}
