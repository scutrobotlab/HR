<template>
  <v-container>
    <v-text-field
      label="Client ID"
      prepend-icon="mdi-lock"
      v-model="model.client_id"
    ></v-text-field>
    <v-text-field
      label="Client Secret"
      prepend-icon="mdi-key"
      v-model="model.client_secret"
    ></v-text-field>
    <v-text-field
      label="面试系统跳转地址"
      prepend-icon="mdi-cat"
      v-model="model.redirect_uri"
    ></v-text-field>
    <v-text-field
      label="统一认证地址"
      prepend-icon="mdi-shield"
      v-model="model.admin_uri"
    ></v-text-field>
    <v-text-field
      label="微信认证地址"
      prepend-icon="mdi-wechat"
      v-model="model.wechat_uri"
    ></v-text-field>
    <v-btn @click="set">保存</v-btn>
  </v-container>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";

export default {
  mixins: [errorMixin],
  data() {
    return {
      model: {
        client_id: "",
        client_secret: "",
        admin_uri: "",
        wechat_uri: "",
        redirect_uri: "",
      },
    };
  },
  mounted() {
    this.errorHandler(this.$store.dispatch("oauth/sudo_get")).then(() => {
      this.model = this.$store.state.oauth;
    });
  },
  methods: {
    set() {
      this.errorHandler(this.$store.dispatch("oauth/set", this.model));
    },
  },
};
</script>

<style></style>
