import { useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import axios from 'axios';
import useAuthStore from '../store/AuthStore';

function GoogleAuthCallback() {
  //   const [auth, setAuth] = useState();
  const location = useLocation();
  const navigate = useNavigate();
  const setIsLoggedIn = useAuthStore((state) => state.setIsLoggedIn);
  useEffect(() => {
    const { search } = location;
    axios
      .get(`https://be.codein.studio/auth/callback/google${search}`, {
        withCredentials: false,
      })
      .then((res) => {
        // setAuth(res?.data.data);
        setIsLoggedIn(true);
        window.localStorage.setItem('key', res?.data?.data?.token);
        window.localStorage.setItem('idUser', res?.data?.data?.user?.id);
        navigate('/');
      })
      .catch((err) => {
        console.log(err);
        navigate('/');
      });
    //eslint-disable-next-line
  }, []);
}

export default GoogleAuthCallback;
