<template>
  <nav
    role="navigation"
    aria-label="main navigation"
    class="navbar is-transparent main-nav"
  >
    <div class="container">
      <div class="navbar-brand">
        <a href="javascript:" class="navbar-item">
          <img :src="require('@/assets/metamask.png')" ref="metamask-logo" />
        </a>
        <div class="navbar-item">
          <span>PPV Events</span>
        </div>
        <a
          role="button"
          id="nav-toggle"
          aria-label="menu"
          aria-expanded="false"
          data-target="navbarCounterNetwork"
          class="navbar-burger burger"
        >
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>
      <div id="navbarCounterNetwork" class="navbar-menu">
        <div class="navbar-end">
          <div class="navbar-item">
            <a href="javascript:" class="navbar-link is-arrowless">
              Info
            </a>
          </div>
          <div class="navbar-item has-dropdown is-hoverable">
            <a href="javascript:" class="navbar-link is-arrowless">
              <span class="icon" style="display: contents;">
                <span>🇬🇧</span>
                &nbsp;
                <span>English </span>
              </span>
            </a>
            <div class="navbar-dropdown">
              <a href="javascript:" class="navbar-item language-item">
                <span class="icon has-text-info">
                  🇺🇸
                </span>
                &nbsp;
                <span>English (United States) </span>
              </a>
              <a href="javascript:" class="navbar-item language-item">
                <span class="icon has-text-info">🇪🇸</span>
                &nbsp;
                <span>Español </span>
              </a>
            </div>
          </div>
          <div class="navbar-item has-button">
            <div class="buttons">
              <router-link
                class="button is-chalk is-outlined"
                :to="{ name: 'login' }"
              >
                Log in
              </router-link>
              <a href="javascript:" class="button is-primary"
                ><strong>Register now</strong></a
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<style lang="scss" scoped src="./header.scss"></style>
