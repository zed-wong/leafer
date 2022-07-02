export const state = () => ({
    vault: localStorage.getItem('data') ? JSON.parse(localStorage.getItem('data')).vault : [],
    user: localStorage.getItem('data') ? JSON.parse(localStorage.getItem('data')).user : {},
    lodata: localStorage.getItem('data') ? JSON.parse(localStorage.getItem('data')) : {},
    priceset: localStorage.getItem("priceset") ? localStorage.getItem('priceset') : [],
    lang: navigator.language ? navigator.language.substring(0, 2) : "en"
})

export const mutations = {
    poll(state, data) {
        state.data = data
        localStorage.setItem('data', JSON.stringify(data))
    },
    updateVaults(state, vaults) {
        state.vault = vaults
    },
    updateUser(state, user) {
        state.user = user
    },
    updateTg(state, { tgID, tgName }) {
        state.user.tg_id = tgID
        state.user.tg_name = tgName
    },
    updateNumber(state, { tp, number }) {
        switch (tp) {
            case "signal":
                state.user.signal_number = number
                break;
            case "phone":
                state.user.phone_number = number
                break;
        }
    },
    updateRatioByID(state, { id, ratio }) {
        let targetIndex = state.vault.findIndex((obj => obj.identity_id == id));
        state.vault[targetIndex].alert_ratio = ratio.toString()
    },
    removeVaultByID(state, id) {
        var removed = state.vault.filter(function (el) { return el.identity_id != id; });
        state.vault = removed
    },
    updatePriceset(state, priceset){
        state.priceset = priceset
        localStorage.setItem('priceset', priceset)
    }
}