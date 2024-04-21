import viteLogo from '../../public/vite.svg'

export const Register = () => {
  return (
    <>
      <header>
        <div className="max-w-7xl flex justify-center">
          <div className='flex'>
            <img src={viteLogo} className="logo" alt="Vite logo" />
            <h1>ファクトリー</h1>
          </div>
          <div>
            <button>ログイン</button>
          </div>
        </div>
      </header>
      <div className="bg-gray">
        <div className='w-2/3 mx-auto py-8'>
          <div>
            <h1>加工のご依頼・ご相談</h1>
          </div>
          <div className="bg-white rounded py-11 px-8">
            <form action="">
              <h2>会社情報</h2>
              <table className='w-full'>
                <tbody>
                  <tr>
                    <th  className='border border-slate-700'>会社名</th>
                    <th  className='border border-slate-700'>
                      <input type="text" />
                    </th>
                  </tr>
                  <tr>
                    <th  className='border border-slate-700'>会社住所</th>
                    <th className='border border-slate-700'>
                      <div>
                        <input type="text" placeholder='000-0000'/>
                      </div>
                      <div>
                        <input type="text" placeholder='東京都'/>
                        <input type="text" placeholder='中央区'/>
                      </div>
                      <div>
                        <input type="text" placeholder='銀座7-11-15'/>
                      </div>
                    </th>
                  </tr>
                  <tr>
                    <th  className='border border-slate-700'>電話番号</th>
                    <th  className='border border-slate-700'>
                      <input type="text" />
                    </th>
                  </tr>
                </tbody>
              </table>
              <h2>個人情報</h2>
              <table className='w-full'>
                <tbody>
                  <tr>
                    <th className='border border-slate-700'>名前</th>
                    <th className='border border-slate-700'>
                      <input type="text" placeholder='山田太郎'/>
                    </th>
                  </tr>
                  <tr>
                    <th className='border border-slate-700'>email</th>
                    <th className='border border-slate-700'>
                        <input type="text" placeholder='yamada_taro@example.com'/>
                    </th>
                  </tr>
                  <tr>
                    <th className='border border-slate-700'>パスワード</th>
                    <th className='border border-slate-700'>
                      <input type="text" placeholder='パスワード'/>
                    </th>
                  </tr>
                </tbody>
              </table>
            </form>
          </div>
        </div>
      </div>
    </>
  )
}
