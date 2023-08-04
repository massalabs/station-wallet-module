import { useEffect, useState } from 'react';

import { Button, toast } from '@massalabs/react-ui-kit';
import { FiPlus } from 'react-icons/fi';
import { useParams } from 'react-router-dom';

import { AssetsImportModal } from './AssetsImportModal';
import { AssetsList } from './AssetsList';
import { AssetsLoading } from './AssetsLoading';
import { useDelete, useResource } from '@/custom/api';
import Intl from '@/i18n/i18n';
import { WalletLayout, MenuItem } from '@/layouts/WalletLayout/WalletLayout';
import { IToken } from '@/models/AccountModel';

function Assets() {
  const [modal, setModal] = useState(false);
  const { nickname } = useParams();
  const { data: tokenArray = [], isLoading: isGetLoading } = useResource<
    IToken[]
  >(`accounts/${nickname}/assets`);

  const {
    mutate: mutableDelete,
    isSuccess: isSuccessDelete,
    isError: isErrorDelete,
    error: errorDelete,
  } = useDelete<IToken[]>(`accounts/${nickname}/assets`);

  useEffect(() => {
    if (isSuccessDelete) {
      toast.success('Token Deleted Successfully');
    } else if (isErrorDelete) {
      console.log(errorDelete);
    }
  }, [isSuccessDelete]);

  return (
    <WalletLayout menuItem={MenuItem.Assets}>
      <div className="bg-secondary p-10 h-fit w-2/3 rounded-lg">
        <div className="flex items-center justify-between w-full mb-6">
          <div> {Intl.t('assets.title')}</div>
          <Button
            customClass="w-fit"
            preIcon={<FiPlus size={114} />}
            onClick={() => {
              setModal(true);
            }}
          >
            {Intl.t('assets.import')}
          </Button>
        </div>
        <div className="flex flex-col w-full h-fit bg-primary rounded-lg gap-4 p-8">
          {isGetLoading ? (
            <AssetsLoading />
          ) : (
            <AssetsList tokenArray={tokenArray} mutableDelete={mutableDelete} />
          )}
        </div>
        {modal && <AssetsImportModal setModal={setModal} />}
      </div>
    </WalletLayout>
  );
}

export default Assets;
