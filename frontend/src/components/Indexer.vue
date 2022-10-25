<template>
  <div>
    <b-modal class="modal" id="modal-see-more" size="lg" :title="text">
      <p class="my-4">Subject: {{ email.subject }}</p>
      <p class="my-4">From: {{ email.from }}</p>
      <p class="my-4">To: {{ email.to }}</p>
      <p class="my-4">{{ email.content }}</p>
    </b-modal>
    <b-navbar toggleable="lg" style="background-color: #5E33FF">
      <b-navbar-brand style="color: aliceblue"><b-icon icon="envelope" style="color: aliceblue"></b-icon> Enron email</b-navbar-brand>
    </b-navbar>
    <b-container class="mt-5">
      <b-row>
        <b-col cols="6">
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
          <b-card border-variant="primary">
            <b-card-title class="title-card">
              {{ email.subject }}
            </b-card-title>        
            <b-card-text>{{ email.date }}</b-card-text>
            <b-link  v-b-modal.modal-see-more v-on:click="setEmailInformation(email)" class="card-link">See more</b-link>
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
  data () {
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
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
  .title-card {
    color: #5E33FF;
    display: -webkit-box; 
    -webkit-box-orient: vertical; 
    -webkit-line-clamp: 1; 
    line-clamp: 1; 
    overflow:hidden;
  }
</style>
  