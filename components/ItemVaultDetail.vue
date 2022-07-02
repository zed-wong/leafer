<template>
  <v-card>
    <v-card-title class="text-h5 lighten-4">
      {{ $t("vault.detail.title") }}
    </v-card-title>

    <v-list class="pa-2">
      <v-list-item v-for="item in details" :key="item.id">
        <v-list-item-title class="font-weight-medium">
          {{item.key}}
        </v-list-item-title>
        <v-list-item-subtitle class="text-center">
          {{item.value}}
        </v-list-item-subtitle>
      </v-list-item>
    </v-list>

    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn text @click="over"> {{ $t("cancel") }} </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  methods: {
    over: over,
  },
  props:["vault"],
  data() {
    return {
      details:[
          { key: this.$t("vault.detail.id"), value: this.vault.identity_id},
          { key: this.$t("vault.detail.type"), value: this.$t("vault.detail.paid")},
          { key: this.$t("vault.detail.addat"), value: this.vault.add_at.split("T")[0]},
          { key: this.$t("vault.detail.endat"), value: this.vault.end_at.split("T")[0]},
      ]
    };
  },
};
function over() {
  this.$emit("close-dialog");
}
</script>