<template>
  <v-container>
    <v-row>
      <v-col cols="12" sm="6" v-for="tc in TimeFrames" :key="tc.key">
        <date-time-picker
          :label="$t(tc.key)"
          v-model="tc.value"
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
    TimeFrames: [
      { key: "apply_form_start", value: null },
      { key: "apply_form_end", value: null },
      { key: "select_time_start", value: null },
      { key: "select_time_end", value: null },
      { key: "publish_result_start", value: null },
    ],
    loading: false,
  }),
  mounted() {
    this.loading = true;
    axios({
      method: "get",
      url: "/api/time-frame",
    })
      .then((response) => {
        response.data.forEach((kv) => {
          var time_frame = this.TimeFrames.find((tc) => tc.key === kv.key);
          if (time_frame) time_frame.value = kv.value;
        });
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
        url: "/api/time-frame",
        data: { TimeFrames: TimeFrames },
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