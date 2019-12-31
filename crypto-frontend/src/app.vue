<template>
  <div class="fullscreen">
    <div id="header">
      <div class="landing">
        <app-header></app-header>
        <div class="container fullscreen">
          <router-view />
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" src="./app_global.scss"></style>

<script lang="ts" src="./app.ts"></script>
