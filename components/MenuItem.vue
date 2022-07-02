<template>
  <v-dialog v-model="dialog">
    <template v-slot:activator="{ on, attrs }">
      <v-list-item>
        <v-list-item-title v-bind="attrs" v-on="on">{{
          item.name
        }}</v-list-item-title>
      </v-list-item>
    </template>
    <item-vault-detail v-if="item.num === 1" v-on:close-dialog="closeDialog" :vault="vault"/>
    <item-update-ratio v-else-if="item.num === 2" v-on:close-dialog="closeDialog" :vault="vault"/>
    <item-renew-vault v-else-if="item.num === 3" v-on:close-dialog="closeDialog" :vault="vault" />
    <item-remove-vault v-else-if="item.num === 4" v-on:close-dialog="closeDialog" :vault="vault" />
  </v-dialog>
</template>

<script>
import VaultDetail from "./ItemVaultDetail.vue"
import UpdateRatio from './ItemUpdateRatio.vue'
import ItemRenewVault from './ItemRenewVault.vue';
import ItemRemoveVault from './ItemRemoveVault.vue';

export default {
  components: { VaultDetail, UpdateRatio, ItemRenewVault, ItemRemoveVault },
  props: ["item", "vault"],
  methods: {closeDialog: closeDialog},
  data(){
    return{
      dialog: false
    }
  }
};
function closeDialog(){
  this.dialog = false;
}
</script>