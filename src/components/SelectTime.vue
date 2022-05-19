<template>
  <div>
    <ErrorAlert v-model="error" />
    <v-alert type="info" text class="text-left" v-if="!loaded">正在加载...</v-alert>
    <v-alert type="info" text prominent class="text-left" v-else-if="submitting"
      >正在提交...</v-alert
    >
    <v-alert type="info" text prominent class="text-left" v-else-if="!saved_time_id"
      >请选择第{{ rank }}志愿“{{ group }}”的面试时间。</v-alert
    >
    <v-alert type="success" text prominent class="text-left" v-else>
      <v-row align="center">
        <v-col class="grow py-0">已选择第{{ rank }}志愿“{{ group }}”的面试时间。</v-col>
        <v-col class="shrink py-0">
          <v-btn outlined color="success" @click="nextStep">下一步</v-btn>
        </v-col>
      </v-row>
    </v-alert>
    <v-tabs centered v-model="selectDate" class="none-prev">
      <v-tab v-for="(time, date) in dateTime" :key="date" :href="`#${date}`">{{
        fDate(date)
      }}</v-tab>
    </v-tabs>
    <v-item-group v-model="time_id" @change="submit">
      <v-tabs-items v-model="selectDate">
        <v-tab-item v-for="(times, date) in dateTime" :key="date" :value="date">
          <v-container>
            <v-row>
              <v-col
                v-for="time in times"
                :class="$vuetify.breakpoint.xs ? 'pa-1' : ''"
                :key="time.id"
                cols="4"
                sm="3"
                md="2"
              >
                <v-item v-slot:default="{ active, toggle }" :value="time.id">
                  <v-card
                    :disabled="time.left_cnt <= 0"
                    :color="
                      active
                        ? 'primary'
                        : time.left_cnt <= 0
                        ? 'blue-grey lighten-1'
                        : 'blue-grey darken-3'
                    "
                    :elevation="time.left_cnt <= 0 ? 0 : 3"
                    class="d-flex align-center"
                    dark
                    :height="$vuetify.breakpoint.xs ? 80 : 100"
                    @click="toggle"
                  >
                    <v-list-item two-line>
                      <v-list-item-content>
                        <v-list-item-title
                          :class="
                            ($vuetify.breakpoint.xs ? 'headline' : 'display-1') +
                            ' flex-grow-1 text-center'
                          "
                          >{{ fTime(time.time) }}</v-list-item-title
                        >
                        <v-scroll-y-transition mode="out-in">
                          <v-list-item-subtitle v-if="active" class="text-center" key="active">
                            <v-icon x-small>mdi-checkbox-marked-circle-outline</v-icon>&nbsp;已选择
                          </v-list-item-subtitle>
                          <v-list-item-subtitle v-else class="text-center" key="unactive"
                            >剩余&nbsp;{{ time.left_cnt }}&nbsp;位</v-list-item-subtitle
                          >
                        </v-scroll-y-transition>
                      </v-list-item-content>
                    </v-list-item>
                  </v-card>
                </v-item>
              </v-col>
            </v-row>
          </v-container>
        </v-tab-item>
      </v-tabs-items>
    </v-item-group>
  </div>
</template>

<script>
import moment from "moment";
import axios from "axios";
import errorMixin from "@/mixins/errorMixin.js";

export default {
  mixins: [errorMixin],
  props: {
    group: String,
    rank: Number,
  },
  data: () => ({
    selectDate: null,
    dateTime: null,
    time_id: null,
    submitting: false,
    saved_time_id: null,
    timer: null,
    loaded: false,
  }),

  methods: {
    overtime(time) {
      return moment(time.date + "T" + time.time).isBefore(moment().subtract(1, "hours"));
    },

    fDate(date) {
      return moment(date).format("MM/DD");
    },
    fTime(time) {
      return moment(time, "HH:mm:ss").format("HH:mm");
    },

    nextStep() {
      this.$emit("next");
    },

    update() {
      var groupBy = function (arr, key) {
        return arr.reduce(function (rv, item) {
          (rv[item[key]] = rv[item[key]] || []).push(item);
          return rv;
        }, {});
      };

      this.errorHandler(
        axios({
          method: "get",
          url: "/api/joinus/time",
          params: { group: this.group, rank: this.rank },
        }).then((response) => {
          this.dateTime = groupBy(response.data.time, "date");
          if (response.data.saved_time) {
            this.time_id = this.saved_time_id = response.data.saved_time.id;
            this.selectDate = response.data.saved_time.date;
          }
          this.loaded = true;
        })
      );
    },
    submit() {
      this.submitting = true;
      this.errorHandler(
        axios({
          method: "post",
          url: "/api/joinus/time",
          data: { time_id: this.time_id, group: this.group },
        })
          .then(() => (this.saved_time_id = this.time_id))
          .finally(() => {
            this.submitting = false;
            this.update();
          })
      );
    },
  },
  mounted() {
    this.update();
    this.timer = setInterval(this.update, 10000);
  },
  destroyed() {
    clearInterval(this.timer);
  },
};
</script>

<style></style>
