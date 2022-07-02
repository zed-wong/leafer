<template>
  <v-sheet class="rounded pa-3 mb-4" :color="bgcolor">
    <v-row class="pt-1 d-flex align-center">
      <v-col cols="6" class="py-0">
        <span class="name-text"> #{{ vault.identity_id }} </span>
        <alert-renew :vault="vault"  v-if="!valid"/>
      </v-col>

      <v-col cols="6" class="py-0 text-right">
        <v-menu>
          <template
            v-slot:activator="{ on, attrs }"
            transition="scale-transition"
          >
            <v-btn icon v-bind="attrs" v-on="on">
              <v-icon>mdi-dots-horizontal</v-icon>
            </v-btn>
          </template>
          <v-list>
            <menu-item
              v-for="item in menuitems"
              :key="item.num"
              :item="item"
              :vault="vault"
            />
          </v-list>
        </v-menu>
      </v-col>

      <v-col cols="6" class="pt-6 text-left d-flex justify-start align-center">
        <v-img :src="vault.avatar" max-width="32px" max-height="32px" />
      </v-col>

      <v-col cols="6" class="pt-6 text-right d-flex justify-end align-center">
        <span class="ratio-text"> {{ ratio }}%</span>
      </v-col>
    </v-row>
  </v-sheet>
</template>

<script>
import AlertRenew from './AlertRenew.vue';
import MenuItem from "./MenuItem.vue";
export default {
  components: { MenuItem, AlertRenew },
  props: ["vault"],
  computed: {
    valid() {
      console.log(this.vault.end_at, isValid(this.vault.end_at))
      return isValid(this.vault.end_at);
    },
    ratio() {
      return isNaN(this.vault.ratio) ? 0.0 : this.vault.ratio;
    },
    bgcolor() {
      return this.valid ? "#ebfced" : "#E0E0E0";
    },
  },
  data() {
    return {
      menuitems: [
        {
          num: 1,
          name: this.$t("vault.menu.item1"),
        },
        {
          num: 2,
          name: this.$t("vault.menu.item2"),
        },
        {
          num: 3,
          name: this.$t("vault.menu.item3"),
        },
        {
          num: 4,
          name: this.$t("vault.menu.item4"),
        },
      ],
    };
  },
  methods: {},
};

function isValid(date) {
  // date = date.split("T")[0];
  // let d = date.split("-");
  // return new Date(d[0], d[1], d[2]) > new Date(new Date().toDateString());
  return new Date(Date.parse(date)) > new Date();
}
</script>
<style lang="scss" scoped>
.name-text {
  font-weight: 500;
  font-size: 14px;
  opacity: 0.5;
}
.ratio-text {
  font-size: 26px;
  font-weight: 700;
}
</style>