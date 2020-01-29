<template>
  <div class="container is-fluid">
    <h1 class="title has-text-chalk">Events</h1>
    <h2 class="subtitle has-text-chalk">
      Currently available events to purchase
    </h2>
    <div class="columns is-centered" v-if="userCanCreate">
      <router-link
        class="button is-success is-clickable"
        :to="{ name: 'events-create' }"
      >
        Create new event
      </router-link>
    </div>
    <div class="columns" v-if="isLoading">
      <b-icon pack="fas" icon="sync-alt" custom-class="fa-spin"> </b-icon>
      Searching for events
    </div>
    <template v-else-if="hasAvailableEvents">
      <div class="columns is-centered">
        <div class="column">Available events</div>
      </div>
      <div
        class="columns is-centered is-size-7"
        v-for="e in availableEvents"
        :key="e.id"
      >
        <div class="column">{{ e.name }}</div>
        <div class="column">{{ e.descryption }}</div>
        <div class="column">{{ e.eventType }}</div>
        <div class="column">{{ e.start }} - {{ e.end }}</div>
      </div>
    </template>
    <template v-else-if="hasSubscribedEventsEvents">
      <div class="columns is-centered">
        <div class="column">Subscribed events</div>
      </div>
      <div
        class="columns is-centered is-size-7"
        v-for="e in subscribedEvents"
        :key="e.id"
      >
        <div class="column">{{ e.name }}</div>
        <div class="column">{{ e.descryption }}</div>
        <div class="column">{{ e.eventType }}</div>
        <div class="column">{{ e.start }} - {{ e.end }}</div>
      </div>
      subscribedEvents
    </template>
    <div class="columns is-centered is-size-7" v-else-if="hasError">
      <p>
        <b-icon pack="fas" icon="exclamation-circle"> </b-icon>
        There was an error retrieving events
      </p>
    </div>
    <div class="columns" v-else>
      <p>
        <b-icon pack="fas" icon="search"> </b-icon>
        There are no available events
      </p>
    </div>
  </div>
</template>

<script lang="ts" src="./events.ts"></script>
