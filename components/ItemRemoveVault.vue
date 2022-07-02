<template>
  <v-card>
    <v-card-title class="text-h5">
      {{ $t("vault.remove.title") }}
    </v-card-title>
    <v-card-text class="py-2 font-weight-medium">
      {{ $t("vault.remove.text") }} #{{ vault.identity_id }}?
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn text @click="over">
        {{ $t("cancel") }}
      </v-btn>
      <v-btn color="primary" text @click="remove(vault.identity_id)">
        {{ $t("vault.remove.remove") }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  methods: {
    over: over,
    remove: remove,
  },
  props: ["vault"],
};
function over() {
  this.$emit("close-dialog");
}
async function remove(id) {
  let token = this.$store.state.user.access_token;
  let header = {
    Authorization: `Bearer ${token}`,
    UserID: this.$store.state.user.user_id,
  };

  try {
    const resp = await this.$axios.delete("/delete/vault", {
      headers: header,
      data: {
        identity_id: id,
      },
    });
    if (resp.status == 200) {
      this.$emit("close-dialog");
      this.$store.commit("removeVaultByID", id);
      this.$router.push("/");
    }
  } catch (err) {
    console.log(err);
  }
}
</script>