<template>
  <div class="container">
    <event-edit-form v-if="!eventNotFound"></event-edit-form>
    <not-found v-else></not-found>
  </div>
</template>

<script lang="ts" src="./event-form.ts"></script>
