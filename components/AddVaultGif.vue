<template>
  <v-dialog v-model="dialog" width="500">
    <template v-slot:activator="{ on, attrs }">
      <span v-bind="attrs" v-on="on" class="blue--text">{{
        $t("vault.add.help")
      }}</span>
    </template>

    <v-card>
      <v-row no-gutters>
        <v-img
          :src="gifpath"
          class="mb-3 align-center"
          :aspect-ratio="2 / 3.5"
          width="200"
        >
        </v-img>
      </v-row>
    </v-card>
  </v-dialog>
</template>

<script>
var ENpath = require("../static/EN.gif");
var ZHpath = require("../static/ZH.gif");
export default {
  data() {
    return {
      dialog: false,
      gifpath: this.$store.state.lang == "en" ? ENpath : ZHpath,
    };
  },
};
</script>