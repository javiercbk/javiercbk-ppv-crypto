<template>
  <div class="container is-fluid">
    <h1 class="title has-text-chalk">Events</h1>
    <h2 class="subtitle has-text-chalk">
      Currently available events to purchase
    </h2>
    <div class="columns" v-if="isLoading">
      <b-icon pack="fas" icon="sync-alt" custom-class="fa-spin"> </b-icon>
      Searching for events
    </div>
    <div class="columns is-centered is-size-7" v-else-if="hasEvents">
      <p>
        <b-icon pack="fas" icon="search"> </b-icon>
        There are no available events
      </p>
    </div>
    <div class="columns is-centered is-size-7" v-else-if="hasEvents">
      <p>
        <b-icon pack="fas" icon="exclamation-circle"> </b-icon>
        There was an error retrieving events
      </p>
    </div>
    <div class="columns" v-else></div>
  </div>
</template>

<script lang="ts" src="./events.ts"></script>
