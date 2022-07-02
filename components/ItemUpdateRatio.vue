<template>
  <v-card>
    <v-card-title class="text-h5 lighten-4">
      {{ $t("vault.update.text") }}
    </v-card-title>
    <v-row no-gutters>
      <v-col cols="10" class="mt-3">
        <span class="pl-6">
          {{ $t("vault.update.current") }}: {{ vault.alert_ratio }} %</span
        >
      </v-col>
      <v-col cols="10">
        <v-card-text>
          <v-slider
            :max="max"
            :min="min"
            v-model="val"
            color="red"
            class="align-center"
            hide-details
          >
            <template v-slot:append>
              <v-text-field
                :max="max"
                :min="min"
                v-model="val"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 55px"
              ></v-text-field>
              <p class="mt-2 pt-0">%</p>
            </template>
          </v-slider>
        </v-card-text>
      </v-col>
    </v-row>
    <v-card-actions>
      <v-spacer />
      <v-btn text @click="over">
        {{ $t("cancel") }}
      </v-btn>
      <v-btn
        color="primary"
        text
        class=""
        @click="submit(vault.identity_id, val)"
      >
        {{ $t("vault.update.btn") }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: "ItemUpdateRatio",
  methods: {
    over,
    submit,
  },
  props: ["vault"],
  data() {
    return {
      val: this.vault.alert_ratio,
      max: 1000,
      min: 160,
      color: "red",
    };
  },
};
function over() {
  this.$emit("close-dialog");
}

async function submit(id, ratio) {
  let token = this.$store.state.user.access_token;
  let config = {
    headers: {
      Authorization: `Bearer ${token}`,
      UserID: this.$store.state.user.user_id,
    },
  };

  try {
    const resp = await this.$axios.put("/update/ratio", {
    identity_id: id,
    ratio: ratio,
  }, config);
    if (resp.status == 200) {
      this.$emit("close-dialog");
      this.$store.commit("updateRatioByID", {id, ratio});
    }
  } catch (err) {
    console.log(err);
  }
}
</script>