<template>
  <v-layout column>
    <vault-overview :vaults="allvaults" class="mb-4" />
    <span
      class="subtitle-1 font-weight-bold pl-1 d-flex align-center py-4 sticky"
    >
      {{ $t("vault.my") }}
    </span>
    <vault-card
      v-for="vault in allvaults"
      :key="vault.identity_id"
      :vault="vault"
    />
    <add-vault />
  </v-layout>
</template>

<script>
import AddVault from "~/components/AddVault.vue";
import VaultCard from "~/components/VaultCard.vue";
import VaultOverview from "../components/VaultOverview.vue";
export default {
  components: { VaultOverview, AddVault, VaultCard },

  computed: {
    allvaults() {
      return this.$store.state.vault;
    },
  },
};
</script>

<style lang="scss" scoped>
.sticky {
  background-color: white;
  position: sticky;
  top: 44px;
  z-index: 14;
}
</style>