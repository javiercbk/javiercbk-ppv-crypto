<template>
  <section class="hero is-fullheight-with-navbar">
    <app-header></app-header>
    <div class="hero-body is-flex-start">
      <router-view />
    </div>
    <app-footer></app-footer>
  </section>
</template>

<style lang="scss" src="./app_global.scss"></style>

<script lang="ts" src="./app.ts"></script>
