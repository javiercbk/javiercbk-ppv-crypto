<template>
  <div class="container">
    <div class="columns is-centered">
      <div class="column is-5-tablet is-4-desktop is-3-widescreen">
        <form action="" class="box" @submit="createEvent" novalidate>
          <div class="field">
            <label for="event-form-name" class="label">Email</label>
            <div class="control has-icons-left">
              <input
                id="event-form-name"
                type="name"
                v-model="name"
                class="input"
                :class="{ 'is-danger': $v.name.$invalid && $v.name.$dirty }"
                :disabled="loadingEvent"
              />
            </div>
          </div>
          <div class="field">
            <label for="event-form-description" class="label"
              >Description</label
            >
            <div class="control has-icons-left">
              <textarea
                id="event-form-description"
                type="description"
                v-model="description"
                class="textarea"
                :class="{
                  'is-danger': $v.description.$invalid && $v.description.$dirty
                }"
                :disabled="loadingEvent"
              ></textarea>
            </div>
          </div>
          <div class="field">
            <label for="event-form-type" class="label">Event type</label>
            <div class="control has-icons-left">
              <input
                id="event-form-type"
                type="text"
                v-model="eventType"
                class="input"
                :class="{
                  'is-danger': $v.eventType.$invalid && $v.eventType.$dirty
                }"
                :disabled="loadingEvent"
              />
            </div>
          </div>
          <div class="field">
            <label for="event-form-start" class="label"
              >Start datetime in UTC</label
            >
            <div class="control has-icons-left">
              <b-datetimepicker
                id="event-form-start"
                rounded
                v-model="start"
                :datepicker="{ showWeekNumber: false }"
                :timepicker="{ enableSeconds: false, hourFormat: '12' }"
                :class="{
                  'is-danger': $v.start.$invalid && $v.start.$dirty
                }"
                :disabled="loadingEvent"
              >
              </b-datetimepicker>
            </div>
          </div>
          <div class="field">
            <label for="event-form-end" class="label"
              >End datetime in UTC</label
            >
            <div class="control has-icons-left">
              <b-datetimepicker
                id="event-form-end"
                rounded
                v-model="end"
                :datepicker="{ showWeekNumber: false }"
                :timepicker="{ enableSeconds: false, hourFormat: '12' }"
                :min-datetime="start"
                :class="{
                  'is-danger': $v.end.$invalid && $v.end.$dirty
                }"
                :disabled="loadingEvent"
              >
              </b-datetimepicker>
            </div>
          </div>
          <div class="field">
            <label for="event-form-btc" class="label">Price BTC</label>
            <div class="control has-icons-left">
              <input
                id="event-form-btc"
                type="number"
                v-model="priceBTC"
                class="input"
                :class="{
                  'is-danger': $v.priceBTC.$invalid && $v.priceBTC.$dirty
                }"
                :disabled="loadingEvent"
              />
            </div>
          </div>
          <div class="field">
            <label for="event-form-xmr" class="label">Price XMR</label>
            <div class="control has-icons-left">
              <input
                id="event-form-xmr"
                type="number"
                v-model="priceXMR"
                class="input"
                :class="{
                  'is-danger': $v.priceXMR.$invalid && $v.priceXMR.$dirty
                }"
                :disabled="loadingEvent"
              />
            </div>
          </div>
          <div class="field">
            <label for="event-form-smart-contract" class="label"
              >ETH smart contract address</label
            >
            <div class="control has-icons-left">
              <input
                id="event-form-smart-contract"
                type="number"
                v-model="ethContractAddr"
                class="input"
                :class="{
                  'is-danger':
                    $v.ethContractAddr.$invalid && $v.ethContractAddr.$dirty
                }"
                :disabled="loadingEvent"
              />
            </div>
          </div>
          <div class="field">
            <button
              class="button is-success"
              :disabled="$v.$invalid && $v.$dirty"
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" src="./event-create-form.ts"></script>
