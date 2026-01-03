import { createContext, useEffect, useState } from 'react';

const AuthContext = createContext({});

export const AuthProvider = ({ children }) => {
  const [auth, setAuth] = useState();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      try {
        setAuth(JSON.parse(storedUser));
      } catch {
        console.log('Failed to parse user from localStorage');
      }
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    if (auth) {
      localStorage.setItem('user', JSON.stringify(auth));
    } else {
      localStorage.removeItem('user');
    }
  }, [auth]);

  return <AuthContext.Provider value={{ auth, setAuth, loading }}>{children}</AuthContext.Provider>;
};

export default AuthContext;
