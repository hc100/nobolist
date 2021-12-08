<template>
  <nav class="flex items-center justify-between flex-wrap bg-teal-500 p-6">
    <div class="flex items-center flex-shrink-0 text-black mr-6">
      <span class="font-semibold text-xl tracking-tight"
        ><a href="/">ﾉﾎﾞﾘｽﾄ</a></span
      >
    </div>
    <div class="block lg:hidden">
      <button
        class="flex items-center px-3 py-2 border rounded text-teal-200 border-teal-400 hover:text-black hover:border-black"
        @click="toggleMenu"
      >
        <svg
          class="fill-current h-3 w-3"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <title>Menu</title>
          <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
        </svg>
      </button>
    </div>
    <div
      v-if="showMenu"
      class="w-full block flex-grow lg:flex lg:items-center lg:w-auto"
    >
      <div class="text-sm lg:flex-grow">
        <a
          href="#responsive-header"
          class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-black mr-4"
        >
          課題検索
        </a>
        <a
          href="#responsive-header"
          class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-black mr-4"
        >
          Examples
        </a>
        <a
          v-if="!isAuthenticated"
          href="/register/"
          class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-black"
        >
          会員登録
        </a>
      </div>
      <div>
        <a
          v-if="isAuthenticated"
          href="/logout/"
          class="inline-block text-sm px-4 py-2 leading-none border rounded text-black border-black hover:border-transparent hover:text-teal-500 hover:bg-black mt-4 lg:mt-0"
          >ログアウト</a
        >
        <a
          v-else
          href="/login/"
          class="inline-block text-sm px-4 py-2 leading-none border rounded text-black border-black hover:border-transparent hover:text-white hover:bg-black mt-4 lg:mt-0"
          >ログイン</a
        >
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'Header',
  props: {},
  data: () => ({
    showMenu: true,
    isAuthenticated: false,
  }),
  apollo: {},
  watch: {},
  mounted() {
    this.isAuthenticated = !!this.$apolloHelpers.getToken()
  },
  created() {},
  methods: {
    toggleMenu() {
      this.showMenu = !this.showMenu
    },
  },
}
</script>
