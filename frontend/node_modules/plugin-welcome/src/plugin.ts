import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/Users';
import Teacher from './components/teacher';
import Login from './components/Login';
import Table from './components/Table';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Table', Table);
    router.registerRoute('/Teacher', Teacher);
    router.registerRoute('/Users',CreateUser);
    router.registerRoute('/', Login);
  },
});