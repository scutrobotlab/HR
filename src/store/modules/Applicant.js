import axios from "axios";

export default {
  namespaced: true,
  state: {
    wechat: null,
    intents: [],
    step: 0,
    model: {
      wechat_id: null,
      name: "",
      gender: "",
      phone: "",
      intents: [],
      parallel: false,
      form: {},
    }
  },
  mutations: {},
  actions: {
    setModel({state}, model) {
      state.model = model;
    },
    getModel({ state }) {
      return axios({
        method: "get",
        url: "/api/joinus/form",
      }).then((response) => {
        if (response.data.applicant) {
          state.model.name = response.data.applicant.name;
          state.model.gender = response.data.applicant.gender;
          state.model.phone = response.data.applicant.phone;
          if (response.data.applicant.intents)
            state.model.intents = response.data.applicant.intents.map((intent) => intent.group);
          state.model.form = response.data.applicant.form;
        }
      })
    },

    wechatInfo({ state }) {
      return axios.get("/api/joinus/wechat").then((response) => {
        state.wechat = response.data;
      });
    },

    wechatLogin({ state }, token) {
      return axios.post("/api/joinus/login/" + token).then((response) => {
        state.wechat = response.data;
      });
    },

    getStep({ state }) {
      return axios.get("/api/joinus/step/").then((response) => {
        state.step = response.data.step;
      });
    },

    getIntents({ state }) {
      return axios({
        method: "get",
        url: "/api/joinus/intents",
        params: { wechat_id: state.wechat.openid },
      }).then((response) => {
        state.intents = [];
        state.intents = response.data.intents.map((intent) => intent.group);
      })
    },
  },
};
