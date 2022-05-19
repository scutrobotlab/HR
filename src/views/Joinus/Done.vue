<template>
  <v-container fluid>
    <ErrorAlert v-model="error" />
    <v-alert type="info" text v-if="step === 1">您还没有提交报名表。</v-alert>
    <v-alert type="info" text v-else-if="step === 2">您还有未选择的面试时间。</v-alert>
    <div v-else class="text-center mt-5">
      <div v-if="body">
        {{ body }}
      </div>
      <div v-else>
        <v-progress-circular indeterminate color="primary"></v-progress-circular>
      </div>
    </div>
    <div class="text-center" v-if="step > 1">
      <v-btn x-large text to="/joinus/exam/main" color="primary">
        各组题库抢先看<br />
        测试面试成功率
      </v-btn>
    </div>
  </v-container>
</template>

<script>
import moment from "moment";
import axios from "axios";
import Handlebars from "handlebars";
import errorMixin from "@/mixins/errorMixin.js";

export default {
  mixins: [errorMixin],
  data: () => ({
    body: "",
    result: null,
  }),
  methods: {
    fDate(date) {
      return moment(date).format("MMMDo");
    },
    fTime(time) {
      return moment(time, "HH:mm:ss").format("HH:mm");
    },
  },
  mounted() {
    this.errorHandler(
      axios({
        method: "get",
        url: "/api/joinus/result",
      }).then((response) => {
        this.result = response.data;
      })
    );
  },
  watch: {
    result(new_val) {
      var data = {
        applicant: this.applicant,
        admitted: this.result.groups,
        times: this.result.times,
      };
      if (new_val.status === "passed") {
        this.errorHandler(
          axios.get("/api/announce/Admitted").then((response) => {
            const template = Handlebars.compile(response.data.body);
            this.body = template(data);
          })
        );
      } else if (new_val.status === "failed") {
        this.errorHandler(
          axios.get("/api/announce/Failed").then((response) => {
            const template = Handlebars.compile(response.data.body);
            this.body = template(data);
          })
        );
      } else {
        this.errorHandler(
          axios.get("/api/announce/HaventPublished").then((response) => {
            const template = Handlebars.compile(response.data.body);
            this.body = template(data);
          })
        );
      }
    },
  },
  computed: {
    applicant() {
      return this.$store.state.applicant.model;
    },
    step() {
      return this.$store.state.applicant.step;
    },
    intents() {
      return this.$store.state.applicant.intents;
    },
  },
};
</script>

<style scoped>
</style>
