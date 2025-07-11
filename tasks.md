# **Doodle Jump風ゲーム 開発タスクリスト (Ebitengine版)**

Go言語のゲームエンジン「Ebitengine」を使い、Doodle Jump風ゲームを制作するためのタスクリストです。各タスクはコード生成エージェントへの指示として利用できるよう、Ebitengineの実装要件を詳細に記述しています。

### **PR \#1: プロジェクトのセットアップとプレイヤーの描画**

**ゴール:** Ebitengineプロジェクトを初期化し、薄い青の背景のウィンドウ内に、白いプレイヤー（矩形）を描画する。

* **タスク詳細:**
    1. Goモジュールを初期化する (go mod init \<module\_name\>)。
    2. main.goファイルを作成する。
    3. プレイヤーを表すPlayer structを定義する。フィールドは以下を含む:
        * X, Y, Width, Height float64
        * Image \*ebiten.Image (プレイヤーのグラフィック用)
    4. ゲーム全体の状態を管理するGame structを定義し、ebiten.Gameインターフェース（Update, Draw, Layoutメソッド）を実装する。Playerをフィールドとして持つ。
    5. main関数でGame structを初期化する。この時、PlayerのImageをebiten.NewImage(width, height)で作成し、白 (color.White) で塗りつぶす。プレイヤーの初期位置 (X, Y) を設定する。
    6. ebiten.SetWindowSizeでウィンドウサイズを(400, 600)に設定し、ebiten.RunGameでゲームループを開始する。
    7. Drawメソッド内で、背景を薄い青 (例: color.RGBA{R: 173, G: 216, B: 230, A: 255}) でクリアする。
    8. ebiten.DrawImageOptionsを使ってプレイヤーの描画位置を指定し、screen.DrawImage(g.player.Image, op)でプレイヤーを描画する。

### **PR \#2: プレイヤーの左右移動**

**ゴール:** キーボードの左右矢印キーでプレイヤーを操作できるようにする。

* **タスク詳細:**
    1. Game structのUpdateメソッド内にロジックを記述する。
    2. ebiten.IsKeyPressed(ebiten.KeyLeft)を使い、左矢印キーが押されているか判定する。押されていればプレイヤーのX座標から定数（例: 5）を減算する。
    3. 同様にebiten.IsKeyPressed(ebiten.KeyRight)で右矢印キーを判定し、X座標に加算する。
    4. プレイヤーが画面外に出ないようにX座標を制限する。Xが0未満になったら0に、ウィンドウ幅 \- player.Widthを超えたらその値に補正する。

### **PR \#3: 重力とジャンプの導入**

**ゴール:** プレイヤーが自然に落下し、画面の底に着くとジャンプするようになる。

* **タスク詳細:**
    1. Player structに、垂直方向の速度 VelocityY float64 を追加する。
    2. Game structに物理定数として Gravity float64 \= 0.5 を定義する。
    3. Updateメソッド内で、毎フレームplayer.VelocityYにGravityを加算する。
    4. プレイヤーのY座標にplayer.VelocityYを加算する。
    5. プレイヤーのY座標がウィンドウ高さ \- player.Heightを超えた場合（画面の底に着いた場合）:
        * Y座標をウィンドウ高さ \- player.Heightに固定する。
        * VelocityYにジャンプ力として-12を設定する。

### **PR \#4: 足場の生成と描画**

**ゴール:** プレイヤーが乗るための足場を複数、ランダムな位置に描画する。

* **タスク詳細:**
    1. 足場を表すPlatform structを定義する（Playerと同様にX, Y, Width, Height, Imageを持つ）。
    2. Game structに、足場のスライス Platforms \[\]\*Platform を追加する。
    3. ゲーム初期化時に、指定した数（例: 5つ）のPlatformインスタンスを生成する。
        * 各PlatformのImageをebiten.NewImageで作成し、緑 (color.RGBA{R: 0, G: 255, B: 0, A: 255}) で塗りつぶす。
        * X座標はrand.Float64()を使い、ウィンドウ幅の範囲内でランダムに決定する。
        * Y座標は100から550の間に均等に配置する。
    4. Drawメソッド内で、Platformsスライスをループし、全ての足場を描画する。

### **PR \#5: プレイヤーと足場の衝突判定**

**ゴール:** プレイヤーが落下中に足場の上に着地すると、ジャンプするようになる。

* **タスク詳細:**
    1. Updateメソッド内に衝突判定ロジックを記述する。
    2. プレイヤーのVelocityYが0より大きい（落下中である）ことを確認する。
    3. Platformsスライスをループし、各足場とプレイヤーの矩形が重なっているか判定する（AABB衝突判定）。
    4. 衝突条件は、プレイヤーの矩形と足場の矩形が重なっていること。
    5. 衝突が検出された場合、プレイヤーのVelocityYにジャンプ力として-12を設定する。

### **PR \#6: 画面の左右ループ**

**ゴール:** プレイヤーが画面の片方の端から消えると、反対側の端から現れるようにする。

* **タスク詳細:**
    1. Updateメソッド内のプレイヤー移動処理を修正する。
    2. プレイヤーのX座標が-player.Widthより小さくなったら、X座標をウィンドウ幅に設定する。
    3. プレイヤーのX座標がウィンドウ幅より大きくなったら、X座標を-player.Widthに設定する。
    4. PR \#2で追加した画面端での移動制限ロジックは削除する。

### **PR \#7: カメラのスクロール処理**

**ゴール:** プレイヤーが画面を上昇するのに合わせて、カメラが追従するように見せる。

* **タスク詳細:**
    1. Updateメソッド内で、プレイヤーのY座標がウィンドウ高さ / 2より小さくなった場合の処理を追加する。
    2. この条件を満たした場合、プレイヤーのY座標を直接変更する代わりに、Platformsスライス内の全ての足場のY座標に、プレイヤーの上昇分 (-player.VelocityY) を加算する。
    3. これにより、プレイヤーは画面中央付近に留まり、足場が下にスクロールするように見える。

### **PR \#8: 足場の再生成**

**ゴール:** ゲームが無限に続くように、画面外に消えた足場を再利用して上部に新しい足場を生成する。

* **タスク詳細:**
    1. カメラのスクロール処理（PR \#7）に追記する。
    2. 足場を下にスクロールさせた後、Platformsスライスをループする。
    3. 足場のY座標がウィンドウ高さを超えた（画面外に出た）場合:
        * その足場のY座標を0 \- platform.Heightに再設定する。
        * X座標をrand.Float64()を使い、再度ランダムに設定する。

### **PR \#9: スコアシステムの実装**

**ゴール:** 到達した高さをスコアとして画面に表示する。

* **タスク詳細:**
    1. Game structにScore intとMaxScore intフィールドを追加する。
    2. カメラがスクロールする（PR \#7の）たびに、Scoreをインクリメントする。
    3. ScoreがMaxScoreを超えたらMaxScore \= Scoreとする。
    4. Drawメソッドの最後に、ebiten/textパッケージを使ってスコアを描画する。
        * golang.org/x/image/font/basicfontのbasicfont.Face7x13などをフォントとして使用する。
        * text.Drawを使い、画面左上にfmt.Sprintf("Score: %d", g.MaxScore)で整形した文字列を描画する。

### **PR \#10: ゲームオーバー処理**

**ゴール:** プレイヤーが画面下に落下したらゲームを終了し、リスタートできるようにする。

* **タスク詳細:**
    1. Game structにGameOver boolフィールドを追加する。
    2. Updateメソッド内で、プレイヤーのY座標がウィンドウ高さを超えたらg.GameOver \= trueに設定する。
    3. Updateメソッドの冒頭でif g.GameOverの分岐を追加する。
        * trueの場合、ebiten.IsKeyPressed(ebiten.KeyEnter)をチェックし、押されたらゲームの状態をリセットするメソッド（例: g.Reset()）を呼び出す。その後はreturnして以降の更新処理をスキップする。
    4. Drawメソッドで、g.GameOverがtrueの場合、ゲーム画面の上に"Game Over"と"Press Enter to Restart"のテキストをtext.Drawで描画する。
    5. Reset()メソッドをGame structに実装する。このメソッドはプレイヤーの位置、速度、スコア、足場の位置を初期状態に戻し、g.GameOverをfalseにする。