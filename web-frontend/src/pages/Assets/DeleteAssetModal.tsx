import { useEffect, useState } from 'react';

import {
  Button,
  Input,
  PopupModal,
  PopupModalContent,
  PopupModalHeader,
  toast,
} from '@massalabs/react-ui-kit';
import { AxiosError } from 'axios';
import { FiTrash2 } from 'react-icons/fi';
import { useParams } from 'react-router-dom';

import {
  DeleteAssetsErrors,
  assetDeleteErrors,
  deleteConfirm,
} from '@/const/assets/assets';
import { useDelete } from '@/custom/api';
import Intl from '@/i18n/i18n';
import { IToken } from '@/models/AccountModel';

export function DeleteAssetModal({ ...props }) {
  const { tokenAddress, setModalOpen, refetch } = props;
  const { nickname } = useParams();

  const [deletePhrase, setDeletePhrase] = useState<string>('');
  const [error, setError] = useState<DeleteAssetsErrors | null>(null);

  const {
    mutate: mutateDelete,
    isSuccess: isSuccessDelete,
    isError: isErrorDelete,
    error: errorDelete,
  } = useDelete<IToken[]>(
    `accounts/${nickname}/assets?assetAddress=${tokenAddress}`,
  );

  const axiosError = errorDelete as AxiosError;
  const deleteErrorStatus = axiosError?.response?.status;

  useEffect(() => {
    if (isSuccessDelete) {
      toast.success(Intl.t('assets.delete.success'));
      setModalOpen(false);
      refetch();
    } else if (isErrorDelete) {
      displayErrors(deleteErrorStatus);
    }
  }, [isSuccessDelete, isErrorDelete]);

  function displayErrors(postStatus: number | undefined) {
    switch (postStatus) {
      case assetDeleteErrors.badRequest:
        toast.error(Intl.t('assets.delete.bad-request'));
        break;
      case assetDeleteErrors.invalidAddress:
        toast.error(Intl.t('assets.delete.invalid-address'));
        break;
      case assetDeleteErrors.serverError:
        toast.error(Intl.t('assets.internal-server-error'));
        break;
      default:
        toast.error(Intl.t('assets.unkown-error'));
        console.log('Unknown Error:', deleteErrorStatus);
    }
  }

  function validate(phrase: string) {
    if (!phrase) {
      setError({ phrase: Intl.t('assets.delete.no-value') });
      return false;
    }

    if (phrase !== deleteConfirm) {
      setError({ phrase: Intl.t('assets.delete.wrong-value') });
      return false;
    }
    return true;
  }

  function confirmDelete() {
    if (!validate(deletePhrase)) return;

    mutateDelete({} as IToken[]);
  }

  return (
    <PopupModal
      customClass="!w-[440px] h-1/2"
      fullMode={true}
      onClose={() => setModalOpen(false)}
    >
      <PopupModalHeader>
        <div className="flex flex-col mb-4">
          <div className="mas-title mb-4">{Intl.t('assets.delete.title')}</div>

          <div className="mas-body2">{Intl.t('assets.delete.subtitle')}</div>
        </div>
      </PopupModalHeader>
      <PopupModalContent>
        <div className="mb-4">
          <Input
            value={deletePhrase}
            name="provider"
            onChange={(e) => setDeletePhrase(e.target.value)}
            error={error?.phrase}
            placeholder={Intl.t('assets.delete.placeholder')}
          />
        </div>
        <div className="flex gap-4 pb-12">
          <Button onClick={() => setModalOpen(false)}>
            {Intl.t('assets.delete.cancel')}
          </Button>
          <Button posIcon={<FiTrash2 />} onClick={() => confirmDelete()}>
            {Intl.t('assets.delete.delete')}
          </Button>
        </div>
      </PopupModalContent>
    </PopupModal>
  );
}
