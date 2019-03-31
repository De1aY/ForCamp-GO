<template>
  <div class="dashboard">
    <div class="settings">
      <div class="settings-header">
        <div class="settings-title">Основная информация</div>
      </div>
      <div class="settings-body">
        <div class="settings-field" @dblclick="switchInputStatus('organization')">
          <div class="settings-label">Название организации</div>
          <input type="text" class="settings-input" placeholder="Лицей Иннополис"
                 v-model="organizationName"
                 :disabled="isOrganizationInputDisabled">
        </div>
        <div class="settings-field" @dblclick="switchInputStatus('period')">
          <div class="settings-label">Название периода</div>
          <input type="text" class="settings-input" placeholder="1 четверть"
                 v-model="periodName"
                 :disabled="isPeriodInputDisabled">
        </div>
        <div class="settings-field" @dblclick="switchInputStatus('participant')">
          <div class="settings-label">Название участника</div>
          <input type="text" class="settings-input" placeholder="Ученик"
                 v-model="participantName"
                 :disabled="isParticipantInputDisabled">
        </div>
        <div class="settings-field" @dblclick="switchInputStatus('team')">
          <div class="settings-label">Название команды</div>
          <input type="text" class="settings-input" placeholder="Класс"
                 v-model="teamName"
                 :disabled="isTeamInputDisabled">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      isTeamInputDisabled: true,
      isPeriodInputDisabled: true,
      isParticipantInputDisabled: true,
      isOrganizationInputDisabled: true,
    };
  },
  computed: {
    organizationName: {
      get() {
        return this.$store.state.global.organizationInfo.organizationName;
      },
      set(value) {
        this.$store.commit('setOrgsetOrganizationName', value);
      },
    },
    periodName: {
      get() {
        return this.$store.state.global.organizationInfo.periodName;
      },
      set(value) {
        this.$store.commit('setOrgsetPeriodName', value);
      },
    },
    participantName: {
      get() {
        return this.$store.state.global.organizationInfo.participantsName;
      },
      set(value) {
        this.$store.commit('setOrgsetParticipantName', value);
      },
    },
    teamName: {
      get() {
        return this.$store.state.global.organizationInfo.teamName;
      },
      set(value) {
        this.$store.commit('setOrgsetTeamName', value);
      },
    },
  },
  methods: {
    switchInputStatus(inputName) {
      switch (inputName) {
        case 'team':
          this.isTeamInputDisabled = !this.isTeamInputDisabled;
          break;
        case 'period':
          this.isPeriodInputDisabled = !this.isPeriodInputDisabled;
          break;
        case 'participant':
          this.isParticipantInputDisabled = !this.isParticipantInputDisabled;
          break;
        case 'organization':
          this.isOrganizationInputDisabled = !this.isOrganizationInputDisabled;
          break;
        default:
          break;
      }
    },
  },
};
</script>

<style lang="scss" scoped>
@import "../../../assets/scss/colors";

.dashboard {
  margin: 60px;

  &-title {
    font-size: 24px;
    font-family: "Fira Sans Condensed", sans-serif;
    font-weight: 500;
  }

  .settings {
    width: 350px;
    min-height: 120px;
    border-radius: 4px;
    background: #fff;

    &-header {
      box-sizing: content-box;
      padding: 0 20px;
      height: 60px;
      display: flex;
      justify-content: flex-start;
      align-items: center;
      border-bottom: 2px solid $neutralLightest;
    }

    &-title {
      color: #000;
      font-size: 20px;
      font-family: "Fira Sans Condensed", sans-serif;
    }

    &-body { padding: 15px 20px; }

    &-field {
      height: 70px;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: flex-start;
    }

    &-label {
      color: $neutralDarker;
      font-size: 15px;
      font-family: "Fira Sans Condensed", sans-serif;
    }

    &-input {
      margin-top: 5px;
      border: none;
      color: #000;
      font-size: 20px;
      font-family: "Open Sans", sans-serif;

      &:disabled {
        cursor: default;
        background: #fff;
      }

    }

  }

}
</style>
