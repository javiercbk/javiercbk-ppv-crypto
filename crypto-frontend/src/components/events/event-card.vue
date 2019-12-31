<template>
  <div class="card">
    <div class="card-image">
      <figure class="image is-4by3">
        <img
          src="https://bulma.io/images/placeholders/1280x960.png"
          :alt="name"
        />
      </figure>
    </div>
    <div class="card-content">
      <div class="media">
        <div class="media-left">
          <figure class="image is-48x48">
            <img
              src="https://bulma.io/images/placeholders/96x96.png"
              :alt="name"
            />
          </figure>
        </div>
        <div class="media-content">
          <p class="title is-4">{{ name }}</p>
          <p class="subtitle is-6">{{ eventType }}</p>
        </div>
      </div>

      <div class="content">
        {{ description }}
        <br />
        <time :datetime="starts">{{ starts }}</time>
        <br />
        <time :datetime="ends">{{ ends }}</time>
        <br />
        <div>
          <span><event-pay :currency="cryptoCurrency.ETH"/></span>
          <span
            ><event-pay :currency="cryptoCurrency.ETH" frontend-payment
          /></span>
          <span><event-pay :currency="cryptoCurrency.BTC"/></span>
          <span><event-pay :currency="cryptoCurrency.XMR"/></span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" src="./event-card.ts"></script>
