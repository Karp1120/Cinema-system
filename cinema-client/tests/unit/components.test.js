import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router'
import Login from '../../src/views/Login.vue'
import Films from '../../src/views/Films.vue'
import Cinemas from '../../src/views/Cinemas.vue'
import Sessions from '../../src/views/Sessions.vue'
import Reports from '../../src/views/Reports.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [{ path: '/', component: Login }],
})

describe('Login.vue', () => {
  it('renders login form', async () => {
    const wrapper = mount(Login, {
      global: { plugins: [router] },
    })
    expect(wrapper.text()).toContain('Вход')
  })
})

describe('Films.vue', () => {
  it('renders film section', () => {
    const wrapper = mount(Films)
    expect(wrapper.text()).toContain('Список фильмов')
  })
})

describe('Cinemas.vue', () => {
  it('renders cinemas title', () => {
    const wrapper = mount(Cinemas)
    expect(wrapper.text()).toContain('Список кинотеатров')
  })

  it('renders form inputs for cinema name', () => {
    const wrapper = mount(Cinemas)
    const input = wrapper.find('input[placeholder="Название"]')
    expect(input.exists()).toBe(true)
  })
})

describe('Sessions.vue', () => {
  it('renders sessions section', () => {
    const wrapper = mount(Sessions)
    expect(wrapper.text()).toContain('Список сеансов')
  })

  it('has a form to create session', () => {
    const wrapper = mount(Sessions)
    expect(wrapper.find('form').exists()).toBe(true)
  })
})

describe('Reports.vue', () => {
  it('renders report sections', () => {
    const wrapper = mount(Reports)
    expect(wrapper.text()).toContain('Репертуар по кинотеатру')
  })

  it('has section for action genre cinemas', () => {
    const wrapper = mount(Reports)
    expect(wrapper.text()).toContain('Кинотеатры по жанру')
  })

  it('has form to get seat info', () => {
    const wrapper = mount(Reports)
    expect(wrapper.text()).toContain('Свободные места')
  })

  it("has form for director's films", () => {
    const wrapper = mount(Reports)
    expect(wrapper.text()).toContain('Фильмы по режиссёру')
  })
})
