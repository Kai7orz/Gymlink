import { defineConfig } from 'vite'

export default defineConfig({
  server: {
    allowedHosts: ['katazuke-balancer-1270398432.ap-northeast-1.elb.amazonaws.com']
  }
})
