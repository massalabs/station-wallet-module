import Banner from '../components/Banner/Banner';
import LandingPage from '../layouts/LandingPage/LandingPage';
import Body from '../components/Body';

export default function CreateAccount() {
  return (
    <LandingPage>
      <div className="">
        <Banner>Hey!</Banner>
        <Body>Select an account</Body>
        <button>Create an account</button>
        <button>Import an existing account</button>
      </div>
    </LandingPage>
  );
}
