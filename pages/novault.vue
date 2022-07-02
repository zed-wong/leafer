<template>
  <v-layout align-center justify-center column>
    <div class="greyscale_1--text text-center body-2 mb-8">
      {{ $t("vault.zero.text") }}
    </div>
    <add-vault />
  </v-layout>
</template>

<script>
import AddVault from '~/components/AddVault.vue';
export default {
  components: { AddVault },
};
</script>