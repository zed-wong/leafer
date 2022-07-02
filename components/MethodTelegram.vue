<template>
  <!-- Telegram -->
  <div>
    <method-header :item="item" :user="user" />
    <v-expansion-panel-content>
      <v-row no-gutters>
        <!-- logged in -->
        <v-col
          cols="12"
          class="text-center pt-3 font-weight-medium"
          v-if="user.tg_id"
        >
          {{ item.content }}: {{ item.value }}
        </v-col>
        <!-- to log in -->
        <v-col cols="12" class="text-center pt-3 font-weight-medium" v-else>
          <div>
            <tg-widget
              mode="callback"
              :telegram-login="botID"
              @callback="CallbackFunction"
            />
          </div>
        </v-col>
      </v-row>
    </v-expansion-panel-content>
  </div>
</template>

<script>
import TgWidget from "./TgWidget.vue";
import MethodHeader from "./MethodHeader.vue";
export default {
  components: { TgWidget, MethodHeader },
  props: ["item", "user"],
  data(){
    return{
      botID: process.env.tg_bot_id,
    }
  },
  methods: {
    async CallbackFunction(usr) {
      let tgID = usr.id;
      let tgName = usr.username;
      let token = this.$store.state.user.access_token;
      let config = {
        headers: {
          Authorization: `Bearer ${token}`,
          UserID: this.$store.state.user.user_id,
        },
      };
      try {
        let resp = await this.$axios.put(
          "/update/tg",
          {
            tg_id: tgID,
            tg_name: tgName,
          },
          config
        );
        if (resp.status == 200) {
          this.$store.commit("updateTg", { tgID, tgName });
          this.$router.push("/methods");
        }
      } catch (err) {
        console.log(err);
      }
    },
  },
};
</script>