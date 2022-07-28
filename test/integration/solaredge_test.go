package integration

import (
	"github.com/elliott-davis/solaredge-go/solaredge"
	"log"
	"os"
	"testing"
	"time"
)

var (
	client *solaredge.Client
)

func init() {
	token := os.Getenv("SOLAREDGE_AUTH_TOKEN")
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
	var id int64 = 2249706

	site, err := client.Site.Details(id)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if site.ID != id {
		t.Fatalf("site ID does not match")
	}
	log.Printf("%+v", site)
}

func TestSiteEnergy(t *testing.T) {
	var id int64 = 2249706

	energies, err := client.Site.Energy(id, solaredge.SiteEnergyRequest{
		TimePeriodRequest: solaredge.TimePeriodRequest{
			StartDate: solaredge.YMDTime(time.Date(2022, 07, 25, 0, 0, 0, 0, time.UTC)),
			EndDate:   solaredge.YMDTime(time.Date(2022, 07, 25, 0, 0, 0, 0, time.UTC)),
		},
		TimeUnit: solaredge.QuarterOfAnHour,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(energies) == 0 {
		log.Fatalf("no energies returned")
	}
	for _, value := range energies {
		if value.Value != nil {
			log.Printf("%s: %f", value.Date, *value.Value)
		} else {
			log.Printf("%s: ---", value.Date)
		}
	}
}

func TestSitesEnergies(t *testing.T) {
	ids := []int64{2249706, 1051426, 2935522}

	energies, err := client.Sites.Energy(ids, &solaredge.SiteEnergyRequest{
		TimePeriodRequest: solaredge.TimePeriodRequest{
			StartDate: solaredge.YMDTime(time.Date(2022, 07, 25, 0, 0, 0, 0, time.UTC)),
			EndDate:   solaredge.YMDTime(time.Date(2022, 07, 26, 0, 0, 0, 0, time.UTC)),
		},
		TimeUnit: solaredge.Hour,
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

func TestSiteEnergyDetails(t *testing.T) {
	ids := []int64{2249706}

	for _, id := range ids {
		details, err := client.Site.EnergyDetails(id, solaredge.SiteEnergyDetailsRequest{
			StartTime: solaredge.DateTime{Time: time.Date(2022, 07, 25, 0, 0, 0, 0, time.UTC)},
			EndTime:   solaredge.DateTime{Time: time.Date(2022, 07, 26, 0, 0, 0, 0, time.UTC)},
			Meters:    nil,
			TimeUnit:  solaredge.QuarterOfAnHour,
		})
		if err != nil {
			return
		}
		//log.Printf("%+v", details)
		for _, meter := range details.Meters {
			log.Println(meter.Type.String())
			for _, value := range meter.Values {
				if value.Value != nil {
					log.Println(value.Date, ":", *value.Value)
				}
			}
		}
	}
}
