<template>
  <v-container>
    <v-row>
      <v-col cols="12" sm="6" v-for="(v, k, i) in TimeFrames" :key="i">
        <date-time-picker
          :label="$t(k)"
          :input="v"
          :loading="loading"
          :disabled="loading"
        />
      </v-col>
    </v-row>
    <v-btn :loading="loading" :disabled="loading" @click="submit">提交</v-btn>
  </v-container>
</template>

<script>
import axios from "axios";
import moment from "moment";
import DateTimePicker from "@/components/DateTimePicker.vue";

export default {
  components: { DateTimePicker },
  data: () => ({
    TimeFrames: {
      form_start: null ,
      form_end: null ,
      time_start: null ,
      time_end: null ,
      done: null ,
    },
    loading: false,
  }),
  mounted() {
    this.loading = true;
    axios({
      method: "get",
      url: "/api/public/time-frame",
    })
      .then((response) => {
        this.TimeFrames = response.data;
      })
      .finally(() => (this.loading = false));
  },
  methods: {
    submit() {
      this.loading = true;
      var TimeFrames = this.TimeFrames;
      TimeFrames = TimeFrames.filter((tc) => tc.value);
      TimeFrames.forEach((tc) => (tc.value = moment(tc.value).format("YYYY/MM/DD HH:mm:ss")));
      axios({
        method: "post",
        url: "/api/admin/time-frame",
        data: TimeFrames,
      })
        .then((response) => {
          window.console.log(response.data);
        })
        .finally(() => (this.loading = false));
    },
  },
};
</script>

<style>
</style>