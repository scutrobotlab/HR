<template>
  <v-container fluid>
    <ErrorAlert critical v-model="error" />
    <v-alert type="info" text>{{ prompt_message }}</v-alert>
    <v-form class="text-center" ref="form" v-model="valid" v-if="form" :disabled="!can_apply">
      <v-text-field
        label="姓名"
        name="name"
        v-model="model.name"
        :rules="[(v) => !!v || '姓名是必填项。']"
        required
      />
      <div>
        <v-subheader class="pl-0">性别</v-subheader>
        <v-btn-toggle v-model="model.gender" group mandatory color="primary" name="gender">
          <v-btn value="male" text :disabled="!can_apply">
            <span>男生</span>
            <v-icon right>mdi-gender-male</v-icon>
          </v-btn>
          <v-btn value="female" text :disabled="!can_apply">
            <span>女生</span>
            <v-icon right>mdi-gender-female</v-icon>
          </v-btn>
        </v-btn-toggle>
      </div>
      <v-text-field
        label="手机号码"
        name="phone"
        v-model="model.phone"
        counter="11"
        :rules="[
          (v) => !!v || '手机号码是必填项。',
          (v) => /^\d{11}$/.test(v) || '手机号码必须为11位数字。',
        ]"
        required
      />
      <v-select
        v-if="intent && intent.parallel"
        class="my-2"
        v-model="model.intents"
        name="intent"
        :items="groups"
        :loading="groups_loading"
        :counter="intent.max > 1 ? intent.max : false"
        :rules="[
          (v) => v.length != 0 || '意向组别是必填项。',
          (v) => v.length <= intent.max || '意向组别的数量不能超过' + intent.max + '。',
          (v) => v.length >= intent.min || '意向组别的数量不能少于' + intent.min + '。',
        ]"
        label="意向组别"
        :multiple="intent.max > 1"
        required
      ></v-select>
      <div v-else-if="intent">
        <v-select
          v-for="i in intent.max"
          :key="'intent' + i"
          v-model="model.intents[i - 1]"
          name="intent"
          :items="groups"
          :loading="groups_loading"
          :label="'第' + i + '志愿' + (i > intent.min ? '（选填）' : '')"
          :required="i > intent.min"
          :clearable="i > intent.min"
          :rules="[model.intents.length === new Set(model.intents).size || '多个志愿不能重复。']"
        >
        </v-select>
      </div>
      <div v-for="item in form" :key="item.id">
        <v-select
          v-if="item.type == 'select'"
          v-model="model.form[item.key]"
          :name="item.key"
          :items="item.options"
          :label="item.name"
          :rules="[(v) => !item.required || !!v || item.name + '是必填项。']"
          :required="item.required"
          :clearable="!item.required"
        ></v-select>
        <v-text-field
          v-else-if="item.type === 'text-field'"
          v-model="model.form[item.key]"
          :name="item.key"
          :label="item.name"
          :rules="[(v) => !item.required || !!v || item.name + '是必填项。']"
          :required="item.required"
        ></v-text-field>
        <v-textarea
          v-else-if="item.type === 'textarea'"
          outlined
          v-model="model.form[item.key]"
          :name="item.key"
          :label="item.name"
          :rules="[(v) => !item.required || !!v || item.name + '是必填项。']"
          :required="item.required"
        ></v-textarea>
      </div>
      <v-btn
        @click="submit"
        :loading="submit_loading"
        :disabled="!valid || !can_apply"
        large
        color="primary"
        class="my-2"
        >提交</v-btn
      >
    </v-form>
    <v-row align="center" justify="center" v-else>
      <v-progress-circular indeterminate color="primary"></v-progress-circular>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import moment from "moment";
import errorMixin from "@/mixins/errorMixin.js";

export default {
  mixins: [errorMixin],
  props: {
    apply_msg: String,
  },
  data: () => ({
    intent: null,
    valid: true,
    model: {
      wechat_id: null,
      name: "",
      gender: "",
      phone: "",
      intents: [],
      parallel: false,
      form: {},
    },
    submit_loading: false,
  }),
  watch: {
    stored_model(new_val) {
      window.console.log(new_val);
      this.model = new_val;
    },
  },
  computed: {
    stored_model() {
      return this.$store.state.applicant.model;
    },
    can_apply() {
      return this.$store.state.allowJoinus.applyForm.status;
    },
    prompt_message() {
      let apply = this.$store.state.allowJoinus.applyForm;
      if (!apply) return "";
      let now = moment();
      if (apply.start === null) return "报名尚未开始，请等候通知。";
      else if (now.isBefore(apply.start))
        return "报名将于" + moment(apply.start).fromNow() + "开始。";
      else {
        if (apply.end === null) return "请填写报名表。";
        else if (now.isBefore(apply.end))
          return "报名表填写将于" + moment(apply.end).fromNow() + "截止。";
        else return "报名表填写已经结束。";
      }
    },
    groups_loading() {
      return this.$store.state.groups.loading;
    },
    groups() {
      return this.$store.state.groups.list;
    },
    form() {
      return this.$store.state.form.form;
    },
  },
  methods: {
    submit() {
      this.submit_loading = true;
      this.model.wechat_id = this.$store.state.applicant.wechat.openid;
      this.errorHandler(
        axios({
          method: "post",
          url: "/api/joinus/apply",
          data: this.model,
        })
          .then(() => {
            this.$store.dispatch("applicant/setModel", this.model);
            this.$emit("next");
          })
          .finally(() => (this.submit_loading = false))
      );
    },
  },
  mounted() {
    this.$store.dispatch("groups/get");
    this.$store.dispatch("allowJoinus/canApplyForm");
    this.$store.dispatch("form/get");
    //TODO 应该放到store
    this.errorHandler(
      axios({
        method: "get",
        url: "/api/form/intent",
      }).then((response) => {
        this.intent = response.data.intent;
        this.model.parallel = response.data.intent.parallel;
      })
    );

    this.model = this.stored_model;
  },
};
</script>

<style scoped>
.v-input {
  margin-top: 8px !important;
  margin-bottom: 8px !important;
}
</style>
