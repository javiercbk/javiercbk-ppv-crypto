<template>
  <div class="container-fluid">
    <div class="columns is-centered">
      <div class="column">
        <form action="" class="box" @submit="createEvent" novalidate>
          <div class="field">
            <label for="event-form-name" class="label">Name</label>
            <div class="control has-icons-left">
              <input
                id="event-form-name"
                type="name"
                v-model="name"
                class="input"
                :class="{ 'is-danger': $v.name.$invalid && $v.name.$dirty }"
                :disabled="!isReady"
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
                :disabled="!isReady"
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
                :disabled="!isReady"
              />
            </div>
          </div>
          <b-field label="Start datetime in UTC">
            <b-datetimepicker
              id="event-form-start"
              rounded
              v-model="start"
              :datepicker="{ showWeekNumber: false }"
              :timepicker="{ enableSeconds: false, hourFormat: '12' }"
              :class="{
                'is-danger': $v.start.$invalid && $v.start.$dirty
              }"
              :datetime-formatter="datetimeFormatter"
              :datetime-parser="datetimeParser"
              :disabled="!isReady"
            >
            </b-datetimepicker>
          </b-field>
          <b-field label="End datetime in UTC">
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
              :datetime-formatter="datetimeFormatter"
              :datetime-parser="datetimeParser"
              :disabled="!isReady"
            ></b-datetimepicker>
          </b-field>
          <label for="event-form-us-dollar" class="label"
            >Set value according to US Dollar</label
          >
          <div class="field has-addons">
            <div class="control">
              <input
                id="event-form-us-dollar"
                type="number"
                v-model="usDollars"
                class="input"
                min="0"
                :disabled="!isReady"
              />
            </div>
            <div class="control">
              <button
                class="button"
                :class="{ 'is-loading': !isReady }"
                @click="estimateCryptoValues"
              >
                Estimate
              </button>
            </div>
          </div>
          <div class="field">
            <label for="event-form-btc" class="label"
              >Price BTC (in satoshi)</label
            >
            <div class="control has-icons-left">
              <input
                id="event-form-btc"
                type="number"
                v-model="priceBTC"
                class="input"
                :class="{
                  'is-danger': $v.priceBTC.$invalid && $v.priceBTC.$dirty
                }"
                :disabled="!isReady"
              />
            </div>
          </div>
          <div class="field">
            <label for="event-form-xmr" class="label"
              >Price XMR (in piconero)</label
            >
            <div class="control has-icons-left">
              <input
                id="event-form-xmr"
                type="number"
                v-model="priceXMR"
                class="input"
                :class="{
                  'is-danger': $v.priceXMR.$invalid && $v.priceXMR.$dirty
                }"
                :disabled="!isReady"
              />
            </div>
          </div>
          <div class="field">
            <label for="event-form-xmr" class="label">Price ETH (in wei)</label>
            <div class="control has-icons-left">
              <input
                id="event-form-eth"
                type="number"
                v-model="priceETH"
                class="input"
                :class="{
                  'is-danger': $v.priceETH.$invalid && $v.priceETH.$dirty
                }"
                :disabled="!isReady"
              />
            </div>
          </div>
          <div class="field" v-if="isEdition">
            <label for="event-form-smart-contract" class="label"
              >ETH smart contract address</label
            >
            <div class="control has-icons-left">
              <input
                id="event-form-smart-contract"
                type="number"
                v-model="ethContractAddr"
                class="input"
                disabled
              />
            </div>
          </div>
          <div class="field">
            <button
              class="button is-success"
              :class="{ 'is-loading': !isReady }"
              :disabled="($v.$invalid && $v.$dirty) || !isReady"
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" src="./event-edit-form.ts"></script>
