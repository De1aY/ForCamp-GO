const state = {
  organizationInfo: {
    teamName: 'Класс',
    periodName: '4 четверть',
    organizationName: 'Школа №1',
    participantsName: 'Ученик',
    selfMarks: false,
    statisticsAnonymization: true,
  },
};

const getters = {};

const actions = {};

const mutations = {
  setOrgsetOrganizationName(state, change) {
    state.organizationInfo.organizationName = change;
  },
  setOrgsetTeamName(state, change) {
    state.organizationInfo.teamName = change;
  },
  setOrgsetParticipantName(state, change) {
    state.organizationInfo.participantsName = change;
  },
  setOrgsetPeriodName(state, change) {
    state.organizationInfo.periodName = change;
  },
  switchOrgsetSelfMarks(state) {
    state.organizationInfo.selfMarks = !state.organizationInfo.selfMarks;
  },
  switchOrgsetStatisticsAnonymization(state) {
    // eslint-disable-next-line max-len
    state.organizationInfo.statisticsAnonymization = !state.organizationInfo.statisticsAnonymization;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
