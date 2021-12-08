<template>
  <div>
    <Nav />
    <header class="bg-white shadow">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold text-gray-900">ログアウト</h1>
      </div>
    </header>
    <main class="pl-8"></main>
  </div>
</template>

<script>
export default {
  apollo: {},
  data() {
    return {}
  },
  head() {
    return {
      title: 'ログアウト',
    }
  },
  computed: {},
  watch: {},
  beforeMount() {},
  mounted() {},
  async created() {
    await this.$apolloHelpers.onLogout()
    this.$router.push({ name: 'index' })
  },
  methods: {},
}
</script>

<style></style>
