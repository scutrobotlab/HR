import axios from "axios";

export default {
  namespaced: true,
  state: {
    id: null,
    applicant: null,
    admits: null,
    scores: null,
  },
  mutations: {},
  actions: {
    get({ state }, params) {
      return axios.get("/api/applicant/info/" + params.appId).then((response) => {
        state.id = params.appId;
        state.applicant = response.data.applicant;
        state.admits = response.data.applicant.admits.map((admit) => admit.group);
      });
    },
  },
};
