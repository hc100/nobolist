<template>
  <div>
    <Header />
    <main>
      <div class="w-full max-w-xs">
        <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">ログアウト</div>
      </div>
      <div class="w-full max-w-xs"></div>
    </main>
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
