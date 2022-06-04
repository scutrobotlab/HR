import axios from "axios";

export default {
  namespaced: true,
  state: {
    wechat: null,
    intents: [],
    step: 0,
    model: {
      open_id: null,
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
    get({ state }) {
      return axios({
        method: "get",
        url: "/api/applicant/status",
      }).then((response) => {
        if (response.status === 204) {
          state.step = 0;
          return
        } else if (response.status === 200) {
          state.step = 1;
          state.model.name = response.data.name;
          state.model.gender = response.data.gender;
          state.model.phone = response.data.phone;
          if (response.data.intents)
            state.model.intents = response.data.intents.map((intent) => intent.group);
          state.model.form = response.data.form;
        }
      })
    },

    wechatInfo({ state }) {
      return axios.get("/api/applicant/wechat").then((response) => {
        state.wechat = response.data;
      });
    },

    wechatLogin({ state }, token) {
      return axios.post("/api/applicant/login/" + token).then((response) => {
        state.wechat = response.data;
      });
    },

  },
};
