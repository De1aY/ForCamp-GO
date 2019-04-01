<template>
  <div class="dashboard">
    <div class="settings">
      <div class="settings-header">
        <div class="settings-title">Основная информация</div>
      </div>
      <div class="settings-body">
        <div class="settings-field" @dblclick="switchInputStatus('organization')">
          <input type="text" id="input-organization_name" class="settings-input"
                 placeholder="Лицей Иннополис"
                 v-model="organizationName"
                 :disabled="isOrganizationInputDisabled"
                 @keydown.enter="switchInputStatus('organization')">
          <label for="input-organization_name" class="settings-label">Название организации</label>
        </div>
        <div class="settings-field" @dblclick="switchInputStatus('period')">
          <input type="text" id="input-period_name" class="settings-input" placeholder="1 четверть"
                 v-model="periodName"
                 :disabled="isPeriodInputDisabled"
                 @keydown.enter="switchInputStatus('period')">
          <label for="input-period_name" class="settings-label">Название периода</label>
        </div>
        <div class="settings-field" @dblclick="switchInputStatus('participant')">
          <input type="text" id="input-participant_name" class="settings-input" placeholder="Ученик"
                 v-model="participantName"
                 :disabled="isParticipantInputDisabled"
                 @keydown.enter="switchInputStatus('participant')">
          <label for="input-participant_name" class="settings-label">Название участника</label>
        </div>
        <div class="settings-field" @dblclick="switchInputStatus('team')">
          <input type="text" id="input-team_name" class="settings-input" placeholder="Класс"
                 v-model="teamName"
                 :disabled="isTeamInputDisabled"
                 @keydown.enter="switchInputStatus('team')">
          <label for="input-team_name" class="settings-label">Название команды</label>
        </div>
      </div>
    </div>
    <div class="settings">
      <div class="settings-header">
        <div class="settings-title">Основные настройки</div>
      </div>
      <div class="settings-body">
        <div class="settings-field">
          <input type="checkbox" id="input-self_marks" class="settings-input settings-input--switch"
                 placeholder="Оценки своей команде"
                 v-model="selfMarks">
          <label for="input-self_marks" class="settings-label">Оценки своей команде</label>
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
      teamNameBackup: '',
      periodNameBackup: '',
      participantNameBackup: '',
      organizationNameBackup: '',
      isTeamInputDisabled: true,
      isPeriodInputDisabled: true,
      isParticipantInputDisabled: true,
      isOrganizationInputDisabled: true,
    };
  },
  computed: {
    teamName: {
      get() {
        return this.$store.state.global.organizationInfo.teamName;
      },
      set(value) {
        this.$store.commit('setOrgsetTeamName', value);
      },
    },
    selfMarks: {
      get() {
        return this.$store.state.global.organizationInfo.selfMarks;
      },
      set() {
        this.$store.commit('switchOrgsetSelfMarks');
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
    organizationName: {
      get() {
        return this.$store.state.global.organizationInfo.organizationName;
      },
      set(value) {
        this.$store.commit('setOrgsetOrganizationName', value);
      },
    },
    isTeamNameChanged() {
      return this.teamNameBackup !== this.teamName;
    },
    isPeriodNameChanged() {
      return this.periodNameBackup !== this.periodName;
    },
    isParticipantNameChanged() {
      return this.periodNameBackup !== this.participantName;
    },
    isOrganizationNameChanged() {
      return this.organizationNameBackup !== this.organizationName;
    },
  },
  methods: {
    switchInputStatus(inputName) {
      switch (inputName) {
        case 'team':
          this.isTeamInputDisabled = !this.isTeamInputDisabled;
          if (!this.isTeamInputDisabled) {
            this.teamNameBackup = this.teamName;
          } else if (this.isTeamNameChanged) {
            // TODO: POST request
          }
          break;
        case 'period':
          this.isPeriodInputDisabled = !this.isPeriodInputDisabled;
          if (!this.isPeriodInputDisabled) {
            this.periodNameBackup = this.periodName;
          } else if (this.isPeriodNameChanged) {
            // TODO: POST request
          }
          break;
        case 'participant':
          this.isParticipantInputDisabled = !this.isParticipantInputDisabled;
          if (!this.isParticipantInputDisabled) {
            this.participantNameBackup = this.participantName;
          } else if (this.isParticipantNameChanged) {
            // TODO: POST request
          }
          break;
        case 'organization':
          this.isOrganizationInputDisabled = !this.isOrganizationInputDisabled;
          if (!this.isOrganizationInputDisabled) {
            this.organizationNameBackup = this.organizationName;
          } else if (this.isOrganizationNameChanged) {
            // TODO: POST request
          }
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
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  justify-content: flex-start;

  > div:not(:first-of-type) { margin-left: 60px; }

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

    &-body { padding: 15px 20px 0 20px; }

    &-field {
      position: relative;
      height: 70px;
      display: grid;
      grid-template-rows: 20px 30px 20px;
      grid-template-columns: auto;
    }

    &-label {
      grid-row-start: 1;
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
        user-select: none;

        + .settings-label:after {
          width: 0;
        }

      }

      + .settings-label:after {
        content: ' ';
        position: absolute;
        bottom: 15px;
        left: 0;
        width: 70%;
        height: .8px;
        background: #7f8c8d;
        transition: width 0.3s ease-in-out;
      }

      &--switch {
        display: none;
        cursor: pointer;

        + .settings-label {

          &:after {
            z-index: 2;
            content: ' ';
            position: absolute;
            bottom: 20px;
            left: 0;
            width: 70px;
            height: 28px;
            border-radius: 15px;
            background: red + 150;
            transition: all 0.3s ease-in-out;
          }

          &:before {
            z-index: 3;
            content: ' ';
            position: absolute;
            bottom: 20px;
            left: 0;
            width: 40px;
            height: 28px;
            border-radius: 15px;
            background: red + 25;
            transition: all 0.3s ease-in-out;
          }

        }

        &:checked {

          + .settings-label {

            &:after {
              background: $primary + 100;
            }

            &:before {
              left: 30px;
              background: $primary;
            }

          }

        }

      }

    }

  }

}
</style>
