import Placeholder from '../../layouts/Placeholder/Placeholder';
import WalletLayout, {
  MenuItem,
} from '../../layouts/WalletLayout/WalletLayout';
export default function Home() {
  return (
    <WalletLayout menuItem={MenuItem.Home}>
      <Placeholder />
    </WalletLayout>
  );
}
