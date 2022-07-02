<template>
  <v-dialog v-model="dialog" width="500">
    <template v-slot:activator="{ on, attrs }">
      <v-btn
        rounded
        x-large
        depressed
        color="black"
        class="mb-10 white--text font-weight-bold"
        v-bind="attrs"
        v-on="on"
      >
        {{ $t("vault.add") }}
      </v-btn>
    </template>
    <v-card>
      <v-card-title class="text-h5 text lighten-2">
        {{ $t("vault.add") }}
      </v-card-title>

      <v-card-text>
        {{ $t("vault.add.text") }} <add-vault-gif/>
      </v-card-text>
      <v-card-text>
        {{ $t("vault.add.text2") }}
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" text @click="dialog = false">
          {{ $t("cancel") }}
        </v-btn>
        <v-btn color="primary" text :href="openLeaf">
          {{ $t("vault.add.btn") }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import AddVaultGif from './AddVaultGif.vue';

export default {
  components: { AddVaultGif },
  data() {
    return {
      dialog: false,
      // gifpath: this.$store.state.lang == "en" ? ENpath : ZHpath,
      openLeaf: "mixin://apps/75f18fe8-b056-46d6-9c48-0214425e58ce?action=open",
    };
  },
};
</script>