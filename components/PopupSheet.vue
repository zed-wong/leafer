<template>
<!-- logged in -->
  <v-sheet height="360px" v-if="user.avatar">
    <v-row no-gutters>
      <v-col cols="12" class="text-right">
        <v-btn class="ma-3" icon>
          <v-icon color="black" @click="off"> mdi-close </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-list class="pa-0 font-weight-medium">
      <v-list-item class="mb-3 pl-2" v-for="(item,i) in items" :key="i" @click="click(item.action)">
        <v-list-item-avatar>
          <v-icon color="black"> {{ item.icon }}</v-icon>
        </v-list-item-avatar>
        <v-list-item-title>
          {{ item.title }}
        </v-list-item-title>
      </v-list-item>
    </v-list>

    <v-row no-gutters>
      <v-col cols="12" class="text-center mt-1">
        <v-btn depressed rounded x-large class="red--text" @click="logout()">
          {{ $t("sheet.disconnect") }}
        </v-btn>
      </v-col>
    </v-row>
  </v-sheet>
</template>

<script>
export default {
  props: ["user"],
  data() {
    return {
      items: [
        {
          icon: "mdi-message-alert",
          title: this.$t("sheet.alert.methods"),
          action:"alertMethods",
        },
        {
          icon: "mdi-help-circle",
          title: this.$t("sheet.help.center"),
          action:"helpCenter",
        },
        {
          icon: "mdi-account-multiple-plus",
          title: this.$t("sheet.join.community"),
          action:"joinCommunity",
        }
      ],
    };
  },
  methods: {
    click(action){
      if (action === "alertMethods"){
        this.$router.push('/methods')
      } else if ( action === "helpCenter"){
        this.$router.push('/help')
      } else if ( action === "joinCommunity"){
        window.location.href = process.env.group_link;
      }
    },
    logout(){
      localStorage.removeItem('data')
      this.$store.commit('updateUser', "{}");
      this.$router.push(`/login`);
      this.off()
    },
    off(){
      this.$emit('off')
    }
  }
};
</script>