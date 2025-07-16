# Homebrew配布とCI/CDのセットアップ手順

このドキュメントでは、helm-releaseツールをHomebrewで配布するためのセットアップ手順と、CI/CDパイプラインの設定方法について説明します。

## 実施した変更

1. **コードの更新**
   - GitHubユーザー名を「meithon」に更新
   - インポートパス、モジュールパス、ビルド設定を更新

2. **CI/CD設定**
   - GitHub Actionsワークフローファイルを作成
   - GoReleaser設定ファイルを追加

3. **Homebrew関連ファイル**
   - Homebrew Formula用のテンプレートを作成
   - homebrew-tapリポジトリ用のファイルを準備

## セットアップ手順

セットアップを簡単に行うために、自動化スクリプト `setup-homebrew-and-release.sh` を用意しました。このスクリプトは以下の手順を自動的に実行します：

1. GitHub上に `meithon/homebrew-tap` リポジトリを作成
2. 準備したHomebrew Formulaファイルをリポジトリにプッシュ
3. GitHub Personal Access Tokenの作成をガイド
4. トークンをhelm-releaseリポジトリのシークレットとして設定
5. CHANGELOGを更新して最初のリリースを準備
6. リリースタグを作成してプッシュ

### スクリプトの実行方法

```bash
# スクリプトを実行
./setup-homebrew-and-release.sh
```

スクリプトは対話形式で実行され、各ステップで必要な情報を入力するよう促します。

### 手動でのセットアップ

スクリプトを使用せずに手動でセットアップする場合は、以下の手順に従ってください：

1. **homebrew-tapリポジトリの作成**
   - GitHubで `meithon/homebrew-tap` という名前の新しいリポジトリを作成
   - `homebrew-tap` ディレクトリの内容をリポジトリにプッシュ

2. **GitHub Personal Access Tokenの作成**
   - GitHubの設定から `repo` スコープを持つトークンを作成
   - helm-releaseリポジトリの「Settings > Secrets and variables > Actions」で `HOMEBREW_TAP_TOKEN` という名前のシークレットとして設定

3. **最初のリリースの作成**
   - CHANGELOGを更新
   - タグを作成してプッシュ：`git tag -a v0.1.0 -m "Release v0.1.0" && git push origin v0.1.0`

## インストール方法

セットアップが完了すると、ユーザーは以下のコマンドでhelm-releaseをインストールできるようになります：

```bash
# タップを追加
brew tap meithon/tap

# helm-releaseをインストール
brew install helm-release

# または一行で
brew install meithon/tap/helm-release
```

## CI/CDパイプラインの動作

1. 新しいタグ（例：`v1.0.0`）がプッシュされると、GitHub Actionsワークフローが起動
2. GoReleaserが複数のプラットフォーム用のバイナリをビルド
3. GitHubリリースが作成され、バイナリがアップロード
4. Homebrew Formulaが自動的に更新

## トラブルシューティング

### SHA256ハッシュの問題

Homebrew Formulaのインストール時に以下のようなSHA256ミスマッチエラーが発生した場合：

```
Error: helm-release: SHA256 mismatch
Expected: REPLACE_WITH_ACTUAL_SHA256_AFTER_FIRST_RELEASE
Actual: cc86e974b6f3e9cf03b9d4e7722228ba5f4b68004b5b4ff8a2d9f304fa66cfd0
```

以下のいずれかの方法で解決できます：

1. **update-homebrew-formula.shスクリプトを使用する**：
   ```bash
   ./update-homebrew-formula.sh
   ```
   このスクリプトは自動的にHomebrew Formulaを正しいSHA256ハッシュで更新します。

2. **手動で更新する**：
   ```bash
   # homebrew-tapリポジトリでhelm-release.rbファイルを編集
   sed -i '' 's/REPLACE_WITH_ACTUAL_SHA256_AFTER_FIRST_RELEASE/cc86e974b6f3e9cf03b9d4e7722228ba5f4b68004b5b4ff8a2d9f304fa66cfd0/' helm-release.rb
   git commit -am "Update SHA256 hash"
   git push
   ```

### リリースワークフローの問題

リリースワークフローが失敗した場合：

1. GitHub Actionsのログでエラーを確認
2. 問題を修正
3. 失敗したタグを削除：`git tag -d v1.0.0 && git push --delete origin v1.0.0`
4. 再度リリースプロセスを実行

詳細な情報は各ドキュメントファイル（HOMEBREW_TAP_SETUP.md、RELEASE_PROCESS.md、HOMEBREW_USAGE.md、CI_CD_DETAILS.md）を参照してください。
