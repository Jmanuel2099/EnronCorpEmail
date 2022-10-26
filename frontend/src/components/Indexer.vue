<template>
  <div>
    <b-modal class="modal" scrollable id="modal-see-more" size="lg" ok-only>
      <h5>Subject:</h5>
      <p class="my-2">{{ email.subject }}</p>
      <h5>From: </h5>
      <p class="my-2">{{ email.from }}</p>
      <h5>To:</h5>
      <p class="my-2">: {{ email.to }}</p>
      <p class="my-4">{{ email.content }}</p>
    </b-modal>

    <b-navbar toggleable="lg" style="background-color: #150550">
      <b-navbar-brand class="title-nav">
        <b-icon icon="envelope" class="icon-nav" />
      </b-navbar-brand>
      <b-collapse id="nav-text-collapse" is-nav>
        <b-navbar-nav >
          <b-nav-text class="title-nav" style="color:aliceblue">Enron Emails</b-nav-text>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>

    <b-container class="mt-4">
      <b-row>
        <b-col cols="8">
          <b-input-group size="md" class="mb-2">
            <b-input-group-prepend is-text>
              <b-icon icon="search"></b-icon>
            </b-input-group-prepend>
            <b-form-input v-model="text" placeholder="Search in emails"></b-form-input>
          </b-input-group>
        </b-col>
        <b-col cols="2">
          <b-button variant="outline-primary" :disabled="text === ''" v-on:click="searchInformation()">Search</b-button>
        </b-col>
      </b-row>

      <b-row>
        <b-col cols="4" v-for="email in emails" :key="email.id" class="mt-2 mb-4">
          <b-card border-variant="primary" style="background-color: #B4C9F7">
            <b-card-title class="title-card">
              {{ email.subject }}
            </b-card-title>
            <b-card-text>{{ email.date }}</b-card-text>
            <b-link v-b-modal.modal-see-more v-on:click="setEmailInformation(email)" class="card-link">See more</b-link>
          </b-card>
        </b-col>
      </b-row>

    </b-container>
  </div>
</template>
  
<script>
import axios from 'axios';
import moment from 'moment';
export default {
  name: 'MyIndexer',
  data() {
    return {
      text: '',
      user: '',
      users: [],
      fields: ['subject', 'from', 'to'],
      emails: [],
      email: {},
      showTable: false,
      showContent: false,
      chargeEmails: false,
    }
  },

  created() {
    this.getUsers();
  },

  methods: {
    async searchInformation() {
      let response = await axios.get(`http://localhost:8000/search/${this.text}`);
      this.showTable = true;
      this.emails = response.data.emails.map(email => {
        return {
          subject: email.subject,
          content: email.content,
          to: email.to,
          date: moment(email.date).format("MMM Do YYYY"),
          from: email.from,
          id: email.message_id
        }
      });
    },

    async getUsers() {
      let response = await axios.get('http://localhost:8000/enron/users/');
      this.users = response.data.users;
    },

    async getUserInformation() {
      this.chargeEmails = true;
      let response = await axios.get(`http://localhost:8000/enron/users/${this.user}`);
      if (response) {
        alert('Se insertaron ' + response.data.emails.record_count + ' registros');
      }
      this.chargeEmails = false;
    },

    setEmailInformation(email) {
      this.email = email;
    }
  }
}
</script>
  
<style>
.title-card {
  color: #150550;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
  line-clamp: 1;
  overflow: hidden;
}

.icon-nav {
  color: aliceblue;
  font-size: 40px;
  margin-left: 30px;
}

.title-nav {
  font-family: Roboto, sans-serif;
  font-weight: bold;
  font-size: 28px;
}
</style>
  