package workers

import (
	"fmt"
	"log"
	"time"
	"sync"
	"gorm.io/gorm"
        "gorm.io/driver/postgres"
	"github.com/shopspring/decimal"
	"github.com/zed-wong/leafer/api"
)

type UpdaterWorker struct{
	db	*gorm.DB
}

func NewUpdaterWorker(dsn string) *UpdaterWorker{
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
                panic(err)
        }
	return &UpdaterWorker{
		db: db,
	}
}

func (uw *UpdaterWorker) Update(){
	vaults := api.GetAllVaults(uw.db)
	for _, vault := range *vaults{
		if isActive(vault.EndAt){
			newVault := api.FetchVaultData(vault.VaultID)
			newVault.UserID = vault.UserID
			newVault.AddAt = vault.AddAt
			newVault.EndAt = vault.EndAt
			newVault.AlertRatio = vault.AlertRatio

			if isTriggered(newVault.Ratio, newVault.AlertRatio) || isTriggered(newVault.NextRatio, newVault.AlertRatio){
				newVault.Triggered = true
			} else {
				newVault.Triggered = false
			}

			if !api.UpdateVault(uw.db, *newVault) {
				log.Printf("#%s update failed", vault.IdentityID)
			}
		}
	}
}

func (uw *UpdaterWorker) Check() []api.VAULT{
	triggeredVaults := []api.VAULT{}
        vaults := api.GetAllVaults(uw.db)
        for _, vault := range *vaults{
                if isActive(vault.EndAt){
                        if vault.Triggered {
                                triggeredVaults = append(triggeredVaults, vault)
                        }
                }
        }
        return triggeredVaults
}

func (uw *UpdaterWorker) Alert(vaults []api.VAULT, mw *MethodWorker) {
	var wg sync.WaitGroup
	if len(vaults) == 0 {
		return
	}
	for _, vault := range vaults{
		user := api.GetUser(uw.db, vault.UserID)
		content := api.GetContentByLang(user.Lang)

		wg.Add(1)
	        data := fmt.Sprintf(content.AlertMsg, vault.IdentityID, vault.Ratio)
		censored := fmt.Sprintf(content.CensoredMsg, RandID(4))
		go mw.AlertAllMethods(uw.db, user, &vault, data, censored, &wg)
	}
	wg.Wait()
}

func (uw *UpdaterWorker)CheckAndAlert(mw *MethodWorker){
	uw.Alert(uw.Check(), mw)
}

func isActive(date string) bool{
	d, _ := time.Parse(api.TIMEFORMAT, date)
	return time.Now().Before(d)
}

func isTriggered(ratio, alertRatio string) bool{
	r, _ := decimal.NewFromString(ratio)
	a, _ := decimal.NewFromString(alertRatio)

	// In case of net error
	if r.Equal(decimal.NewFromInt(0)){
		return false
	}
	if a.GreaterThan(r){
		return true
	}
	return false
}
