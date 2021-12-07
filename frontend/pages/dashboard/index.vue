<template>
  <div>
    <Header />
    <main>
      <div class="w-full max-w-xs">
        <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">
          ダッシュボード
        </div>
      </div>
      <div class="w-full max-w-xs"></div>
    </main>
  </div>
</template>

<script>
export default {
  apollo: {},
  middleware: ['isAuth'],
  data: () => ({}),
  head() {
    return {
      title: 'ダッシュボード',
    }
  },
  watch: {},
  mounted() {},
}
</script>
