<template>
  <div class="container">
    <div class="columns is-centered">
      <div class="column is-5-tablet is-4-desktop is-3-widescreen">
        <form action="" class="box" @submit="login" novalidate>
          <div class="field">
            <label for="" class="label">Email</label>
            <div class="control has-icons-left">
              <input
                type="email"
                v-model="email"
                class="input"
                :class="{ 'is-danger': $v.email.$invalid && $v.email.$dirty }"
              />
              <span class="icon is-small is-left">
                <b-icon pack="fas" icon="envelope"> </b-icon>
              </span>
            </div>
          </div>
          <div class="field">
            <label for="" class="label">Password</label>
            <div class="control has-icons-left">
              <input
                type="password"
                v-model="password"
                class="input"
                :class="{
                  'is-danger': $v.password.$invalid && $v.email.$invalid
                }"
              />
              <span class="icon is-small is-left">
                <b-icon pack="fas" icon="lock"> </b-icon>
              </span>
            </div>
          </div>
          <div class="field">
            <button
              class="button is-success"
              :disabled="isRequesting || ($v.$invalid && $v.$dirty)"
            >
              Login
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" src="./login-form.ts"></script>
