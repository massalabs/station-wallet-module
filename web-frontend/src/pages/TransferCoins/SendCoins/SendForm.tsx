import { Balance, Button, Input } from '@massalabs/react-ui-kit';
import { FiArrowUpRight, FiPlus } from 'react-icons/fi';
import Intl from '../../../i18n/i18n';
import Modal from './Advanced';
import ContactList from './ContactList';

export function SendForm({ ...props }) {
  const {
    amount,
    account,
    formattedBalance,
    recipient,
    error,
    setErrorAdvanced,
    errorAdvanced,
    fees,
    modal,
    setModal,
    modalAccounts,
    setModalAccounts,
    handleModalAccounts,
    handleConfirm,
    handleFees,
    setRecipient,
    handleSubmit,
    handleChange,
    SendPercentage,
  } = props;

  const modalArgsAdvanced = {
    fees,
    modal,
    error,
    setErrorAdvanced,
    errorAdvanced,
    setModal,
    handleConfirm,
    handleFees,
  };
  const modalArgsAccounts = {
    modalAccounts,
    setModalAccounts,
    handleModalAccounts,
    setRecipient,
    account,
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        {/* Balance Section */}
        <p className="mas-subtitle mb-5">
          {Intl.t('send-coins.account-balance')}
        </p>
        <Balance customClass="mb-5" amount={formattedBalance} />
        <div className="flex flex-row justify-between w-full pb-3.5 ">
          <p className="mas-body2"> {Intl.t('send-coins.send-action')} </p>
          <p className="mas-body2">
            {Intl.t('send-coins.available-balance')} <u>{formattedBalance}</u>
          </p>
        </div>
        <div className="pb-3.5">
          <Input
            placeholder={'Amount to send'}
            value={amount}
            name="amount"
            onChange={(e) => handleChange(e)}
            error={error?.amount}
          />
        </div>
        <div className="flex flex-row-reverse">
          <ul className="flex flex-row mas-body2">
            <SendPercentage percentage={25} />
            <SendPercentage percentage={50} />
            <SendPercentage percentage={75} />
            <SendPercentage percentage={100} />
          </ul>
        </div>
        <p className="pb-3.5 mas-body2">{Intl.t('send-coins.recipient')}</p>
        <div className="pb-3.5">
          <Input
            placeholder={'Recipient'}
            value={recipient}
            onChange={(e) => setRecipient(e.target.value)}
            name="recipient"
            error={error?.address}
          />
        </div>
        <div className="flex flex-row-reverse pb-3.5">
          <p
            className="hover:cursor-pointer"
            onClick={() => setModalAccounts(!modalAccounts)}
          >
            <u className="mas-body2">
              {Intl.t('send-coins.transfer-between-acc')}
            </u>
          </p>
        </div>
        {/* Button Section */}
        <div className="flex flex-col w-full gap-3.5">
          <Button
            onClick={() => setModal(!modal)}
            variant={'secondary'}
            posIcon={<FiPlus />}
          >
            {Intl.t('send-coins.advanced')}
          </Button>

          <div>
            <Button type="submit" posIcon={<FiArrowUpRight />}>
              {Intl.t('send-coins.send')}
            </Button>
          </div>
        </div>
      </form>
      <div>
        <div>
          {modal ? (
            <Modal {...modalArgsAdvanced} />
          ) : modalAccounts ? (
            <ContactList {...modalArgsAccounts} />
          ) : null}
        </div>
      </div>
    </div>
  );
}
