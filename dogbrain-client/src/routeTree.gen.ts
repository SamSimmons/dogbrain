/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as AuthImport } from './routes/_auth'
import { Route as AnonImport } from './routes/_anon'
import { Route as AnonIndexImport } from './routes/_anon.index'
import { Route as AuthDashboardImport } from './routes/_auth.dashboard'
import { Route as AnonSignUpImport } from './routes/_anon.sign-up'
import { Route as AnonLoginImport } from './routes/_anon.login'
import { Route as AnonCheckEmailImport } from './routes/_anon.check-email'
import { Route as AnonVerifyTokenImport } from './routes/_anon.verify.$token'

// Create Virtual Routes

const AnonAboutLazyImport = createFileRoute('/_anon/about')()

// Create/Update Routes

const AuthRoute = AuthImport.update({
  id: '/_auth',
  getParentRoute: () => rootRoute,
} as any)

const AnonRoute = AnonImport.update({
  id: '/_anon',
  getParentRoute: () => rootRoute,
} as any)

const AnonIndexRoute = AnonIndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => AnonRoute,
} as any)

const AnonAboutLazyRoute = AnonAboutLazyImport.update({
  id: '/about',
  path: '/about',
  getParentRoute: () => AnonRoute,
} as any).lazy(() => import('./routes/_anon.about.lazy').then((d) => d.Route))

const AuthDashboardRoute = AuthDashboardImport.update({
  id: '/dashboard',
  path: '/dashboard',
  getParentRoute: () => AuthRoute,
} as any)

const AnonSignUpRoute = AnonSignUpImport.update({
  id: '/sign-up',
  path: '/sign-up',
  getParentRoute: () => AnonRoute,
} as any)

const AnonLoginRoute = AnonLoginImport.update({
  id: '/login',
  path: '/login',
  getParentRoute: () => AnonRoute,
} as any)

const AnonCheckEmailRoute = AnonCheckEmailImport.update({
  id: '/check-email',
  path: '/check-email',
  getParentRoute: () => AnonRoute,
} as any)

const AnonVerifyTokenRoute = AnonVerifyTokenImport.update({
  id: '/verify/$token',
  path: '/verify/$token',
  getParentRoute: () => AnonRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/_anon': {
      id: '/_anon'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof AnonImport
      parentRoute: typeof rootRoute
    }
    '/_auth': {
      id: '/_auth'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof AuthImport
      parentRoute: typeof rootRoute
    }
    '/_anon/check-email': {
      id: '/_anon/check-email'
      path: '/check-email'
      fullPath: '/check-email'
      preLoaderRoute: typeof AnonCheckEmailImport
      parentRoute: typeof AnonImport
    }
    '/_anon/login': {
      id: '/_anon/login'
      path: '/login'
      fullPath: '/login'
      preLoaderRoute: typeof AnonLoginImport
      parentRoute: typeof AnonImport
    }
    '/_anon/sign-up': {
      id: '/_anon/sign-up'
      path: '/sign-up'
      fullPath: '/sign-up'
      preLoaderRoute: typeof AnonSignUpImport
      parentRoute: typeof AnonImport
    }
    '/_auth/dashboard': {
      id: '/_auth/dashboard'
      path: '/dashboard'
      fullPath: '/dashboard'
      preLoaderRoute: typeof AuthDashboardImport
      parentRoute: typeof AuthImport
    }
    '/_anon/about': {
      id: '/_anon/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AnonAboutLazyImport
      parentRoute: typeof AnonImport
    }
    '/_anon/': {
      id: '/_anon/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof AnonIndexImport
      parentRoute: typeof AnonImport
    }
    '/_anon/verify/$token': {
      id: '/_anon/verify/$token'
      path: '/verify/$token'
      fullPath: '/verify/$token'
      preLoaderRoute: typeof AnonVerifyTokenImport
      parentRoute: typeof AnonImport
    }
  }
}

// Create and export the route tree

interface AnonRouteChildren {
  AnonCheckEmailRoute: typeof AnonCheckEmailRoute
  AnonLoginRoute: typeof AnonLoginRoute
  AnonSignUpRoute: typeof AnonSignUpRoute
  AnonAboutLazyRoute: typeof AnonAboutLazyRoute
  AnonIndexRoute: typeof AnonIndexRoute
  AnonVerifyTokenRoute: typeof AnonVerifyTokenRoute
}

const AnonRouteChildren: AnonRouteChildren = {
  AnonCheckEmailRoute: AnonCheckEmailRoute,
  AnonLoginRoute: AnonLoginRoute,
  AnonSignUpRoute: AnonSignUpRoute,
  AnonAboutLazyRoute: AnonAboutLazyRoute,
  AnonIndexRoute: AnonIndexRoute,
  AnonVerifyTokenRoute: AnonVerifyTokenRoute,
}

const AnonRouteWithChildren = AnonRoute._addFileChildren(AnonRouteChildren)

interface AuthRouteChildren {
  AuthDashboardRoute: typeof AuthDashboardRoute
}

const AuthRouteChildren: AuthRouteChildren = {
  AuthDashboardRoute: AuthDashboardRoute,
}

const AuthRouteWithChildren = AuthRoute._addFileChildren(AuthRouteChildren)

export interface FileRoutesByFullPath {
  '': typeof AuthRouteWithChildren
  '/check-email': typeof AnonCheckEmailRoute
  '/login': typeof AnonLoginRoute
  '/sign-up': typeof AnonSignUpRoute
  '/dashboard': typeof AuthDashboardRoute
  '/about': typeof AnonAboutLazyRoute
  '/': typeof AnonIndexRoute
  '/verify/$token': typeof AnonVerifyTokenRoute
}

export interface FileRoutesByTo {
  '': typeof AuthRouteWithChildren
  '/check-email': typeof AnonCheckEmailRoute
  '/login': typeof AnonLoginRoute
  '/sign-up': typeof AnonSignUpRoute
  '/dashboard': typeof AuthDashboardRoute
  '/about': typeof AnonAboutLazyRoute
  '/': typeof AnonIndexRoute
  '/verify/$token': typeof AnonVerifyTokenRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/_anon': typeof AnonRouteWithChildren
  '/_auth': typeof AuthRouteWithChildren
  '/_anon/check-email': typeof AnonCheckEmailRoute
  '/_anon/login': typeof AnonLoginRoute
  '/_anon/sign-up': typeof AnonSignUpRoute
  '/_auth/dashboard': typeof AuthDashboardRoute
  '/_anon/about': typeof AnonAboutLazyRoute
  '/_anon/': typeof AnonIndexRoute
  '/_anon/verify/$token': typeof AnonVerifyTokenRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths:
    | ''
    | '/check-email'
    | '/login'
    | '/sign-up'
    | '/dashboard'
    | '/about'
    | '/'
    | '/verify/$token'
  fileRoutesByTo: FileRoutesByTo
  to:
    | ''
    | '/check-email'
    | '/login'
    | '/sign-up'
    | '/dashboard'
    | '/about'
    | '/'
    | '/verify/$token'
  id:
    | '__root__'
    | '/_anon'
    | '/_auth'
    | '/_anon/check-email'
    | '/_anon/login'
    | '/_anon/sign-up'
    | '/_auth/dashboard'
    | '/_anon/about'
    | '/_anon/'
    | '/_anon/verify/$token'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  AnonRoute: typeof AnonRouteWithChildren
  AuthRoute: typeof AuthRouteWithChildren
}

const rootRouteChildren: RootRouteChildren = {
  AnonRoute: AnonRouteWithChildren,
  AuthRoute: AuthRouteWithChildren,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* prettier-ignore-end */

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/_anon",
        "/_auth"
      ]
    },
    "/_anon": {
      "filePath": "_anon.tsx",
      "children": [
        "/_anon/check-email",
        "/_anon/login",
        "/_anon/sign-up",
        "/_anon/about",
        "/_anon/",
        "/_anon/verify/$token"
      ]
    },
    "/_auth": {
      "filePath": "_auth.tsx",
      "children": [
        "/_auth/dashboard"
      ]
    },
    "/_anon/check-email": {
      "filePath": "_anon.check-email.tsx",
      "parent": "/_anon"
    },
    "/_anon/login": {
      "filePath": "_anon.login.tsx",
      "parent": "/_anon"
    },
    "/_anon/sign-up": {
      "filePath": "_anon.sign-up.tsx",
      "parent": "/_anon"
    },
    "/_auth/dashboard": {
      "filePath": "_auth.dashboard.tsx",
      "parent": "/_auth"
    },
    "/_anon/about": {
      "filePath": "_anon.about.lazy.tsx",
      "parent": "/_anon"
    },
    "/_anon/": {
      "filePath": "_anon.index.tsx",
      "parent": "/_anon"
    },
    "/_anon/verify/$token": {
      "filePath": "_anon.verify.$token.tsx",
      "parent": "/_anon"
    }
  }
}
ROUTE_MANIFEST_END */
